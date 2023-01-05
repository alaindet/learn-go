package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	r := httprouter.New()
	v := "/v1.0"

	Get(r, v+"/healthcheck", app.healthcheckHandler)
	Post(r, v+"/movies", app.moviesCreateHandler)
	Get(r, v+"/movies/:id", app.moviesShowHandler)

	return r
}

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
