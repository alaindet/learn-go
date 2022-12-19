package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	// Static content serving
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippets/view", app.snippetView)
	mux.HandleFunc("/snippets/create", app.snippetCreate)

	// Middleware
	return app.logRequest(
		secureHeaders(
			mux,
		),
	)
}
