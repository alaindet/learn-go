package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	r := httprouter.New()

	// Custom 404
	r.NotFound = http.HandlerFunc(app.customNotFound)

	// Static content serving
	fileServer := http.FileServer(http.Dir(app.config.staticPath + "/"))
	fileServerHandler := http.StripPrefix("/static", fileServer)
	r.Handler(http.MethodGet, "/static/*filepath", fileServerHandler)

	// Routes
	r.HandlerFunc(http.MethodGet, "/", app.home)
	r.HandlerFunc(http.MethodGet, "/snippets/view/:id", app.snippetView)
	r.HandlerFunc(http.MethodGet, "/snippets/new", app.snippetCreateForm)
	r.HandlerFunc(http.MethodPost, "/snippets", app.snippetCreate)

	// Add global middleware
	globalMiddleware := alice.New(
		app.recoverPanic,
		app.logRequest,
		secureHeaders,
	)

	return globalMiddleware.Then(r)
}
