package main

import (
	"fmt"
	"net/http"
)

func (app *application) signUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}
