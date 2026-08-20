package main

import (
	"app/internal/helpers"
	"net/http"
)

func handleHealthcheck(app *application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data := map[string]string{
			"status":      "available",
			"environment": app.config.env,
			"version":     version,
		}

		err := helpers.WriteJSON(w, http.StatusOK, data, nil)

		if err != nil {
			app.logger.Error(err.Error())
			errMessage := "The server could not process your request"
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})
}
