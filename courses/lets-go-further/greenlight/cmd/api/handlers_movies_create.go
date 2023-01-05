package main

import (
	"fmt"
	"net/http"
)

func (app *application) moviesCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create movie...")
}
