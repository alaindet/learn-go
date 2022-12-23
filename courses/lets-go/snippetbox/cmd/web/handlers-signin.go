package main

import (
	"fmt"
	"net/http"
)

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sign in an existing user")
}
