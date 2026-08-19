package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/v1/healthcheck", handleHealthcheck(app))
	router.Handler(http.MethodPost, "/v1/movies", handleCreateMovie())
	router.Handler(http.MethodGet, "/v1/movies/:id", handleGetMovie(app))

	return router
}
