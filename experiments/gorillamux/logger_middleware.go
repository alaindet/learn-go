package main

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Execute inner http handler regularly
			inner.ServeHTTP(w, r)

			end := time.Since(start)

			log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				end,
			)
		},
	)
}
