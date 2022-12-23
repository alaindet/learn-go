package main

import (
	"fmt"
	"net/http"
)

func (app *application) signOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sign out a logged user")
}
