package main

import (
	"context"
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

// Checks if the user ID from session exists on the database
func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userId := app.sessionManager.GetInt(r.Context(), sessionKeyUserId)

		// No value stored in session, user is not authenticated, leave it as it is
		if userId == 0 {
			next.ServeHTTP(w, r)
			return
		}

		exists, err := app.users.Exists(userId)
		if err != nil {
			app.serverError(w, err)
			return
		}

		if exists {
			ctx := context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
