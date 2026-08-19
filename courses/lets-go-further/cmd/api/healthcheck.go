package main

import (
	"fmt"
	"net/http"
)

func handleHealthcheck(app *application) http.Handler {
	jsonTemplate := `{"status": "available", "environment": %q, "version": %q}`

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonData := fmt.Sprintf(jsonTemplate, app.config.env, version)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonData))
	})
}
