package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {

	avgRequestsPerSecond := rate.Limit(app.config.rateLimit.avg)
	maxRequestsPerSecond := app.config.rateLimit.max

	limiter := rate.NewLimiter(avgRequestsPerSecond, maxRequestsPerSecond)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			app.rateLimitExceededResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
