package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	v := "/v1" // TODO: Move?
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Metrics
	Get(router, v+"/healthcheck", app.healthcheckHandler)

	// Movies
	Get(router, v+"/movies", app.moviesListHandler)
	Post(router, v+"/movies", app.moviesCreateHandler)
	Get(router, v+"/movies/:id", app.moviesShowHandler)
	Patch(router, v+"/movies/:id", app.moviesUpdateHandler)
	Delete(router, v+"/movies/:id", app.moviesDeleteHandler)

	// Middleware
	var handler http.Handler

	if app.config.rateLimiter.enabled {
		handler = app.rateLimit(router)
	}
	handler = app.recoverPanic(handler)

	return handler
}
