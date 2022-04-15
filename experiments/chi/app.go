package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type config struct {
	port string
}

type app struct {
	config *config
	router *chi.Mux
}

func NewApp() *app {
	app := &app{}
	app.init()
	return app
}

func (a *app) init() {
	a.initConfig()
	a.initRouter()
}

func (a *app) initConfig() {
	a.config = &config{
		port: "4000", // TODO: Move to .env
	}
}

func (a *app) initRouter() {
	a.router = chi.NewRouter()
}

func (a *app) SetRouterMiddleware(middlewareFn func(r *chi.Mux)) {
	middlewareFn(a.router)
}

func (a *app) SetRoutes(routesFn func(r *chi.Mux)) {
	routesFn(a.router)
}

func (a *app) SetRouteNotFound(h http.HandlerFunc) {
	a.router.NotFound(h)
}

func (a *app) SetMethodNotAllowed(h http.HandlerFunc) {
	a.router.MethodNotAllowed(h)
}

func (a *app) Start() {
	fmt.Printf("Starting server on port %s\n", a.config.port)
	addr := fmt.Sprintf(":%s", a.config.port)
	http.ListenAndServe(addr, a.router)
}
