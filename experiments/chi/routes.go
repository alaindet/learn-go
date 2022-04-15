package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func registerMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	// Add global middleware here...
}

func registerRoutes(r *chi.Mux) {
	r.Route("/todos", todosRoutes)
	r.Route("/secret", secretRoutes)
	// Add routes here...
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte("Route not found!")
	w.Write(response)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte("Method not allowed!")
	w.Write(response)
}
