package middle

import (
	"base/pkg/cache"
	"base/pkg/config"
	"base/pkg/constant"
	"base/pkg/router"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

// Ratelimit Function : Simple Ratelimit With All Request Per Second For A API
func RateLimit(next http.Handler) http.Handler {
	requestPerSecond, _ := strconv.Atoi(config.Config.GetString("REQUEST_PER_SECOND"))
	requestBurst, _ := strconv.Atoi(config.Config.GetString("REQUEST_BURST"))

	limiter := rate.NewLimiter(rate.Limit(requestPerSecond), requestBurst)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			router.ResponseTooManyRequests(w, constant.TooManyRequests, "")
			return
		}
	})
}

// Ratelimit Function : Set Ratelimit With MaxRequest In A Duration For Specific IP When Make Request
// Block This IP For A Time You Defined
func RateLimitByIP(next http.Handler) http.Handler {
	maxRate, _ := strconv.Atoi(config.Config.GetString("REQUEST_MAX"))
	durationSecond, _ := strconv.Atoi(config.Config.GetString("DURATIONS_SECOND"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// YOU YOU BLOCK IP ADDRESS WHEN THIS IP REQUEST LIMITED

		_, _, isUnderLimit := cache.Redis.Limiter.Allow(r.RemoteAddr, int64(maxRate), time.Duration(durationSecond)*time.Second)
		if !isUnderLimit {
			// Append This IP Into Blocked IP Adderss Slice
			// If You Want To Block This Ip
			// In this Case I don't
			router.ResponseTooManyRequests(w, constant.TooManyRequests, "")
			return

		} else {
			next.ServeHTTP(w, r)
		}
	})
}
