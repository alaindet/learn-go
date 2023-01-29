package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	v := "/v1.0"
	r := httprouter.New()

	r.NotFound = http.HandlerFunc(app.notFoundResponse)
	r.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Metrics
	Get(r, v+"/healthcheck", app.healthcheckHandler)

	// Movies
	Post(r, v+"/movies", app.moviesCreateHandler)
	Get(r, v+"/movies/:id", app.moviesShowHandler)
	Put(r, v+"/movies/:id", app.moviesUpdateHandler)
	Delete(r, v+"/movies/:id", app.moviesDeleteHandler)

	return r
}

// TODO: Move?
func Get(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodGet, path, handler)
}

func Post(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPost, path, handler)
}

func Put(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPut, path, handler)
}

func Delete(r *httprouter.Router, path string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodDelete, path, handler)
}
