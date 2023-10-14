package messages

import (
	"base/pkg/middle"
	"base/service/messages/controller"

	"github.com/go-chi/chi"
)

// UserServiceRoute Create
var MessageServiceRoute = chi.NewRouter()

// Init package with sub-router for account service
func init() {

	MessageServiceRoute.Group(func(r chi.Router) {
		// These APIs USE RateLimit, You Need Config Cache Redis To Run
		// Use Rate Limit Global: Count All Request
		MessageServiceRoute.With(middle.RateLimit).Get("/all", controller.GetAllMessages)

		// Anyone Can Create A Post
		// Use Rate Limit For IP
		MessageServiceRoute.With(middle.RateLimit).Post("/new", controller.CreateMessage)

	})
}
