package main

import (
	"fmt"
	"net/http"
)

func handleHealthcheck(app *application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "status: available")
		fmt.Fprintf(w, "environment: %s\n", app.config.env)
		fmt.Fprintf(w, "version: %s\n", version)
	})
}
