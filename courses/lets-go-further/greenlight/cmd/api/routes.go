package main

import (
	"fmt"
	"net/http"

	h "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	"greenlight/cmd/api/core/middleware"
	monitoring "greenlight/cmd/api/features/monitoring/handlers"
	movies "greenlight/cmd/api/features/movies/handlers"

	"github.com/julienschmidt/httprouter"
)

func Routes(app *core.Application, prefix string) http.Handler {

	router := httprouter.New()
	path := createPathPrefixer(prefix)

	// Standard response
	router.NotFound = http.HandlerFunc(app.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.MethodNotAllowedResponse)

	// Feature: monitoring
	h.Get(router, path("healthcheck"), monitoring.HealthcheckHandler(app))

	// Feature: movies
	h.Get(router, path("movies"), movies.ListHandler(app))
	h.Post(router, path("movies"), movies.CreateHandler(app))
	h.Get(router, path("movies/:id"), movies.ReadHandler(app))
	h.Patch(router, path("movies/:id"), movies.UpdateHandler(app))
	h.Delete(router, path("movies/:id"), movies.DeleteHandler(app))

	// Feature: users
	// ...

	// Middleware
	var handler http.Handler

	// TODO
	if app.Config.RateLimiter.Enabled {
		handler = middleware.RateLimiter(app, router)
	}
	handler = middleware.RecoverPanic(app, handler)

	return handler
}

func createPathPrefixer(prefix string) func(string) string {
	return func(path string) string {
		return fmt.Sprintf("%s/%s", prefix, path)
	}
}
