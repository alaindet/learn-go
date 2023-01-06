package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"status":      "available",
		"environment": app.config.env,
		"version":     app.version,
		"char":        'a',
		"b64":         []byte("This is a string"),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
