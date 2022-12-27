package main

import (
	"fmt"
	"net/http"
)

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !app.isAuthenticated(r) {
			url := r.URL.RequestURI()
			message := fmt.Sprintf("You are not authorized to access \"%s\".", url)
			app.sessionManager.Put(r.Context(), sessionKeyFlash, message)
			http.Redirect(w, r, "/users/signin", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
