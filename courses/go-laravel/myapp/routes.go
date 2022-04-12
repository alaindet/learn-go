package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {

	// ------------------------------------------------------------------------
	//
	// Middleware
	//
	// ------------------------------------------------------------------------

	// Add middleware here...

	// ------------------------------------------------------------------------
	//
	// Routes
	//
	// ------------------------------------------------------------------------

	a.App.Routes.Get("/", a.Handlers.Home)

	// Add routes here...

	// ------------------------------------------------------------------------
	//
	// Serve assets from /public
	//
	// ------------------------------------------------------------------------

	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
