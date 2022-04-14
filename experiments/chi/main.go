package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Init
	r := chi.NewRouter()
	app := NewApp()

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	// TODO: Move handler
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Bootstrap
	addr := fmt.Sprintf(":%s", app.config.port)
	http.ListenAndServe(addr, r)
}
