package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func secretRoutes(r chi.Router) {
	r.With(authMiddleware).Get("/", secretHandler)

	// Equivalent?
	// r.Route("/", func(r chi.Router) {
	// 	r.Use(authMiddleware)
	// 	r.Get("/", secretHandler)
	// })
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := r.URL.Query().Get("password")

		if password == "mellon" {
			next.ServeHTTP(w, r)
			return
		}

		// http.Error(w, http.StatusText(401), 401)
		http.Error(w, "You shall not pass!", 401)
	})
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte("Welcome, my friend.")
	w.Write(response)
}
