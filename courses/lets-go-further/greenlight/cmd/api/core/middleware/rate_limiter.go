package middleware

import (
	"greenlight/cmd/api/core"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	rateLimitClearFrequency = 1 * time.Minute
	rateLimitClearThreshold = 3 * time.Minute
)

func RateLimiter(app *core.Application, next http.Handler) http.Handler {

	if !app.Config.RateLimiter.Enabled {
		return next
	}

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// Clear the map frequently with a separate goroutine
	go func(
		frequency time.Duration,
		removeLimitersOlderThan time.Duration,
	) {
		for {
			time.Sleep(frequency)
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > removeLimitersOlderThan {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}(
		rateLimitClearFrequency,
		rateLimitClearThreshold,
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		mu.Lock()
		_, found := clients[ip]

		// Create a new rate limiter for this IP
		if !found {
			rps := rate.Limit(app.Config.RateLimiter.Rps)
			max := app.Config.RateLimiter.Max
			limiter := rate.NewLimiter(rps, max)
			lastSeen := time.Now()
			clients[ip] = &client{limiter, lastSeen}
		}

		// Update rate limiter and check if client can access
		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			app.RateLimitExceededResponse(w, r)
			return
		}

		mu.Unlock()
		next.ServeHTTP(w, r)
	})
}
