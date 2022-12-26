package main

import (
	"net/http"
)

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.sessionManager.Exists(r.Context(), sessionKeyUserId)
}
