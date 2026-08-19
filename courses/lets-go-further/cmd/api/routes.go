package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/v1/healthcheck", handleHealthcheck(app))
	router.Handler(http.MethodPost, "/v1/movies", handleCreateMovie())
	router.Handler(http.MethodGet, "/v1/movies/:id", handleGetMovie())

	return router
}

// TODO: Move
func handleCreateMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Create movie")
	})
}

// TODO: Move
func handleGetMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Get movie")
	})
}
