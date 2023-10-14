package search

import (
	"base/service/search/consumer"
	"base/service/search/controller"

	"github.com/go-chi/chi"
)

// SearchServiceSubRoute Create
var SearchServiceSubRoute = chi.NewRouter()

// Init package with sub-router for account service
func init() {

	go func() {
		consumer.ConsumeMessage()
	}()

	SearchServiceSubRoute.Group(func(r chi.Router) {
		// Router Here
		SearchServiceSubRoute.Get("/{key_word}", controller.SearchWithKeyWord)

	})
}
