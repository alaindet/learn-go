package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

func TestRoutesExist(t *testing.T) {
	testRoutes := testApp.routes()

	chiRoutes := testRoutes.(chi.Router)

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	findRoute := func(
		method, foundRoute string,
		handler http.Handler,
		middleware ...func(http.Handler) http.Handler,
	) error {
		if route == foundRoute {
			found = true
		}
		return nil
	}

	_ = chi.Walk(routes, findRoute)

	if !found {
		t.Errorf("did not find %s in registered routes", route)
	}
}
