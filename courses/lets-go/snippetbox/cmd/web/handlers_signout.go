package main

import (
	"net/http"
)

func (app *application) signOut(w http.ResponseWriter, r *http.Request) {

	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Remove(r.Context(), sessionKeyUserId)
	app.sessionManager.Put(r.Context(), sessionKeyFlash, "You signed out.")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
