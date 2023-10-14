/*
Package router: Provides a global router for handling HTTP requests.
Package Functionality: This package allows you to create a global router, set middleware for the router, and configure responses and handlers for HTTP requests.
File router.go: Contains the implementation for creating the global router and a health check function for health-related requests.
Author: MinhDan <nguyenmd.works@gmail.com>
*/
package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"base/pkg/cache"
	"base/pkg/config"
	"base/pkg/constant"
	"base/pkg/db"
	"base/pkg/elastic"
	"base/pkg/log"

	"github.com/rs/cors"
)

// RouterBasePath Variable
var RouterBasePath string

// Router Variable
var Router *chi.Mux

// Initialize Function in Router
func init() {

	// Initialize Router
	Router = chi.NewRouter()
	RouterBasePath = config.Config.GetString("ROUTER_BASE_PATH")

	// Create a CORS middleware handler
	c := cors.New(cors.Options{
		// Use this to allow specific origin hosts
		// AllowedOrigins:   []string{"https://example.com"}

		// You can set specific origins or use "*" to allow any
		// AllowedOrigins:   []string{"*"},

		AllowedOrigins:   []string{config.Config.GetString("CORS_ALLOWED_ORIGIN")},
		AllowedMethods:   []string{config.Config.GetString("CORS_ALLOWED_METHOD")},
		AllowedHeaders:   []string{config.Config.GetString("CORS_ALLOWED_HEADER")},
		AllowCredentials: false,
	})

	// Set Router Middleware By Chi
	Router.Use(middleware.RealIP)
	Router.Use(c.Handler)

	// Set Router Handler
	Router.NotFound(handlerNotFound)
	Router.MethodNotAllowed(handlerMethodNotAllowed)
	Router.Get("/favicon.ico", handlerFavIcon)

}

// HealthCheck Function: Check Database, Redis...
// To Sure Service Is Running Good
func HealthCheck(w http.ResponseWriter) {
	// Check Database Connections
	if len(config.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(config.Config.GetString("DB_DRIVER")) {
		case "mysql":
			err := db.MySQL.Ping()
			if err != nil {
				log.Println(log.LogLevelFatal, "health-check", err.Error())
				ResponseInternalError(w, constant.DatabaseConnectionFail, err.Error())
				return
			}
		case "postgres":
			err := db.PSQL.Ping()
			if err != nil {
				log.Println(log.LogLevelError, "health-check", err.Error())
				ResponseInternalError(w, constant.DatabaseConnectionFail, err.Error())
				return
			}
		case "mongo":
			err := db.MongoSession.Ping()
			if err != nil {
				log.Println(log.LogLevelError, "health-check", err.Error())
				ResponseInternalError(w, constant.DatabaseConnectionFail, err.Error())
				return
			}
		}
	}

	// Check Cache Connections : TURN OFF CACHE DRIVER
	if len(config.Config.GetString("REMOTE_CACHE_DRIVER")) != 0 {
		switch strings.ToLower(config.Config.GetString("REMOTE_CACHE_DRIVER")) {
		case "redis":
			// Ping To Check Connection
			_, err := cache.Redis.Ping()

			if err != nil {
				log.Println(log.LogLevelError, "health-check", err.Error())
				ResponseInternalError(w, constant.CacheConnectionFail, err.Error())
				return
			}
		}
	}
	if config.Config.GetString("ENABLE_ELASTIC_SEARCH") == "true" {
		// Ping To Check Connection
		err := elastic.ES.Ping()
		if err != nil {
			log.Println(log.LogLevelError, "health-check", err.Error())
			ResponseInternalError(w, constant.ElasticConnectionFail, err.Error())
			return
		}
	}

	// Return Success response
	ServiceName := config.Config.GetString("SERVER_NAME")
	ResponseSuccess(w, "200", fmt.Sprintf("%s is running good", ServiceName))
}
