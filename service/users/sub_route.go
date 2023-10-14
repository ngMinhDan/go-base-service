package users

import (
	"base/pkg/auth"
	"base/pkg/middle"
	"base/service/users/controller"

	"github.com/go-chi/chi"
)

// UserServiceRoute Create
var UserServiceSubRoute = chi.NewRouter()

// Init package with sub-router for account service
func init() {

	UserServiceSubRoute.Group(func(r chi.Router) {

		// Authentication
		// CRUD Signup, SignIn, SignOut Router - Don't need JWT
		UserServiceSubRoute.Post("/signup", controller.Signup)
		UserServiceSubRoute.Post("/signin", controller.Signin)
		UserServiceSubRoute.Post("/signout", controller.Signout)

		// Authentication and Authorization
		// Access Resource With Role Members
		UserServiceSubRoute.With(auth.JWT).Patch("/change-password", controller.ChangePassword)
		UserServiceSubRoute.With(auth.JWT).Get("/profile/current-profile", controller.GetCurrentProfile)
		UserServiceSubRoute.With(auth.JWT).Patch("/profile/current-profile", controller.UpdateCurrentProfile)

		// Access Resource With Role Admin
		UserServiceSubRoute.With(auth.JWT).Get("/users", controller.GetAllProfile)
		UserServiceSubRoute.With(auth.JWT).Patch("/users/update-role/{id}", controller.UpdateRoleMember)

		// Block Ip Address
		UserServiceSubRoute.With(auth.JWT).Post("/block", controller.BlockIp)

		// Use Middleware To Check This IP In Blocked List
		// You Need * RUN REDIS * For This API
		UserServiceSubRoute.With(middle.IsBlocked).Get("/blocked", controller.CheckBlockedIp)

	})
}
