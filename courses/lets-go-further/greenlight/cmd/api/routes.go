package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	v := "/v1.0" // TODO: Move?
	r := httprouter.New()

	r.NotFound = http.HandlerFunc(app.notFoundResponse)
	r.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Metrics
	Get(r, v+"/healthcheck", app.healthcheckHandler)

	// Movies
	Get(r, v+"/movies", app.moviesListHandler)
	Post(r, v+"/movies", app.moviesCreateHandler)
	Get(r, v+"/movies/:id", app.moviesShowHandler)
	Patch(r, v+"/movies/:id", app.moviesUpdateHandler)
	Delete(r, v+"/movies/:id", app.moviesDeleteHandler)

	return r
}
