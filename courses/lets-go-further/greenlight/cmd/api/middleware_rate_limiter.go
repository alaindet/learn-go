package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {

	// TODO: Make them configurable
	avgRequestsPerSecond := rate.Limit(2)
	maxRequestsPerSecond := 4

	limiter := rate.NewLimiter(avgRequestsPerSecond, maxRequestsPerSecond)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			app.rateLimitExceededResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
