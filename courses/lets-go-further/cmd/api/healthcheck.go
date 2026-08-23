package main

import (
	"net/http"

	"app/internal/helpers"
)

func handleHealthcheck(app *application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{
			"status": "available",
			"systemInfo": map[string]string{
				"environment": app.config.env,
				"version":     version,
			},
		}

		err := helpers.WriteJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.httpErr.InternalServerError(w, r, err)
		}
	})
}
