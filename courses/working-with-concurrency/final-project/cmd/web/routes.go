package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)
	// Add global middleware here...

	// Routes
	mux.Get("/", app.HomePage)
	// Add global routes here...

	return mux
}
