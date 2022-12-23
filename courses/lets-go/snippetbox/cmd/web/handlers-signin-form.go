package main

import (
	"fmt"
	"net/http"
)

func (app *application) signInForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a HTML form for signing in an existing user...")
}
