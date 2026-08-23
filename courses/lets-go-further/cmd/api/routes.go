package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Custom handlers for 404 and 405 HTTP errors
	router.NotFound = http.HandlerFunc(app.handleRouteNotFound)
	router.MethodNotAllowed = http.HandlerFunc(app.handleRouteMethodNotAllowed)

	// Application routes
	router.Handler(http.MethodGet, "/v1/healthcheck", handleHealthcheck(app))
	router.Handler(http.MethodPost, "/v1/movies", handleCreateMovie(app))
	router.Handler(http.MethodGet, "/v1/movies/:id", handleGetMovie(app))

	return app.recoverPanic(router)
}
