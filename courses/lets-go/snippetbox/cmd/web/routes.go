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

	// Routes-only middleware: this chain only applies to application routes,
	// not to static assets routes
	routesMiddleware := alice.New(
		app.sessionManager.LoadAndSave,
	)

	// Routes
	Get(r, "/", routesMiddleware.ThenFunc(app.home))
	Get(r, "/snippets/view/:id", routesMiddleware.ThenFunc(app.snippetView))
	Get(r, "/snippets/new", routesMiddleware.ThenFunc(app.snippetCreateForm))
	Post(r, "/snippets", routesMiddleware.ThenFunc(app.snippetCreate))
	Get(r, "/users/signup", routesMiddleware.ThenFunc(app.signUpForm))
	Post(r, "/users/signup", routesMiddleware.ThenFunc(app.signUp))
	Get(r, "/users/signin", routesMiddleware.ThenFunc(app.signInForm))
	Post(r, "/users/signin", routesMiddleware.ThenFunc(app.signIn))
	Post(r, "/users/signout", routesMiddleware.ThenFunc(app.signOut))

	// Add global middleware
	globalMiddleware := alice.New(
		app.recoverPanic,
		app.logRequest,
		secureHeaders,
	)

	return globalMiddleware.Then(r)
}

func Get(r *httprouter.Router, path string, handler http.Handler) {
	r.Handler(http.MethodGet, path, handler)
}

func Post(r *httprouter.Router, path string, handler http.Handler) {
	r.Handler(http.MethodPost, path, handler)
}
