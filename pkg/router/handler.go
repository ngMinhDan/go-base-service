/*
Package router: Provides a global router for handling HTTP requests.
Package Functionality: This package allows you to create a global router, set middleware for the router, and configure responses and handlers for HTTP requests.
File handler.go: Provide base handler function
Author: MinhDan <nguyenmd.works@gmail.com>
*/
package router

import (
	"net/http"

	"base/pkg/log"
)

// HandlerNotFound Function
func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println(log.LogLevelDebug, "http-access", "not found method "+r.Method+" at URI "+r.RequestURI)
	ResponseNotFound(w, "", "not found method "+r.Method+" at URI"+r.RequestURI)
}

// HandlerMethodNotAllowed Function
func handlerMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Println(log.LogLevelDebug, "http-access", "not allowed method "+r.Method+" at URI "+r.RequestURI)
	ResponseMethodNotAllowed(w, "", "not allowed method "+r.Method+" at URI "+r.RequestURI)
}

// HandlerFavIcon Function
func handlerFavIcon(w http.ResponseWriter, r *http.Request) {
	ResponseNoContent(w)
}
