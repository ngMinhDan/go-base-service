package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"base/pkg/cache"
	"base/pkg/config"
	"base/pkg/db"
	"base/pkg/router"
	"base/pkg/server"

	"base/service"
)

// Server Variable
var svr *server.Server

// Init Function
func init() {
	// Set Go Log Flags
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Load All Routes
	service.LoadRoutes()

	// Initialize Server
	svr = server.NewServer(router.Router)
}

// Main Function
func main() {
	// Starting Server
	svr.Start()

	sig := make(chan os.Signal, 1)
	// Notify Any Signal to OS Signal Channel

	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	// Return OS Signal Channel
	// As Exit Sign
	<-sig
	// Log Break Line
	fmt.Println("")

	// Stopping Server
	defer svr.Stop()

	// Close Any Database Connections
	if len(config.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(config.Config.GetString("DB_DRIVER")) {
		case "postgres":
			log.Println("Stoped connection postgres ...")
			defer db.PSQL.Close()
		case "mysql":
			log.Println("Stoped connection mysql ...")
			defer db.MySQL.Close()
		case "mongo":
			log.Println("Stoped connection mongo ...")
			defer db.MongoSession.Close()
		}
	}

	if strings.ToLower(config.Config.GetString("ENABLE_CACHE_API")) == "true" {
		if len(config.Config.GetString("REMOTE_CACHE_DRIVER")) != 0 {
			switch strings.ToLower(config.Config.GetString("REMOTE_CACHE_DRIVER")) {
			case "redis":
				log.Println("Stoped connection redis ...")
				defer cache.Redis.Close()
			}
		}
	}
}
