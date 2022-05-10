package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, r := range routes {

		// TODO: Refactor
		// Apply global middleware via decorator
		handler := LoggerMiddleware(r.Handler, r.Name)

		router.Methods(r.Method).Path(r.Path).Name(r.Name).Handler(handler)
	}

	return router
}
