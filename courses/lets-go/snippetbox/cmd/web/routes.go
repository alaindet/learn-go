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

	// Setup route-specific middleware
	baseMiddleware := alice.New(
		app.sessionManager.LoadAndSave,
		noSurf,
	)

	baseRoute := func(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
		return baseMiddleware.ThenFunc(handler)
	}

	protectedMiddleware := baseMiddleware.Append(
		app.requireAuthentication,
	)

	protectedRoute := func(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
		return protectedMiddleware.ThenFunc(handler)
	}

	// Routes
	Get(r, "/", baseRoute(app.home))
	Get(r, "/snippets/view/:id", baseRoute(app.snippetView))
	Get(r, "/snippets/new", protectedRoute(app.snippetCreateForm))
	Post(r, "/snippets", protectedRoute(app.snippetCreate))
	Get(r, "/users/signup", baseRoute(app.signUpForm))
	Post(r, "/users/signup", baseRoute(app.signUp))
	Get(r, "/users/signin", baseRoute(app.signInForm))
	Post(r, "/users/signin", baseRoute(app.signIn))
	Post(r, "/users/signout", protectedRoute(app.signOut))

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
