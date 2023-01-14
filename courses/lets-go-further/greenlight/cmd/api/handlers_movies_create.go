package main

import (
	"fmt"
	"net/http"
)

// POST /movies
func (app *application) moviesCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create movie...")
}
