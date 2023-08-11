package main

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {

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
		1*time.Minute,
		3*time.Minute,
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {
			app.internalServerErrorResponse(w, r, err)
			return
		}

		mu.Lock()
		_, found := clients[ip]

		// Create a new rate limiter for this IP
		if !found {
			rps := rate.Limit(app.config.rateLimiter.rps)
			max := app.config.rateLimiter.max
			limiter := rate.NewLimiter(rps, max)
			lastSeen := time.Now()
			clients[ip] = &client{limiter, lastSeen}
		}

		// Update rate limiter and check if client can access
		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			app.rateLimitExceededResponse(w, r)
			return
		}

		mu.Unlock()
		next.ServeHTTP(w, r)
	})
}
