package health

import (
	"base/pkg/router"
	"net/http"
)

// GetIndex Function to Show API Information and Health
func GetIndex(w http.ResponseWriter, r *http.Request) {
	router.HealthCheck(w)
}
