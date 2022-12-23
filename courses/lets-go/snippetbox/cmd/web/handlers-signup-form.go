package main

import (
	"fmt"
	"net/http"
)

func (app *application) signUpForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a HTML form for signing up a new user...")
}
