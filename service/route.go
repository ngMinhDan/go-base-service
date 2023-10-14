package service

import (
	"base/pkg/router"
	"base/service/health"
	"base/service/messages"
	"base/service/search"
	"base/service/users"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {

	// Set Endpoint For Check Health For Service
	router.Router.Get(router.RouterBasePath+"/", health.GetIndex)

	// Mount Users Service's Route To Main Route
	router.Router.Mount(router.RouterBasePath+"/", users.UserServiceSubRoute)

	// Mount Message Service's Route To Main Route
	router.Router.Mount(router.RouterBasePath+"/messages", messages.MessageServiceRoute)

	// Mount Search Service's Route To Main Route
	router.Router.Mount(router.RouterBasePath+"/search", search.SearchServiceSubRoute)

}
