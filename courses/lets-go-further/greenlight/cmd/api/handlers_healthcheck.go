package main

import (
	"fmt"
	"net/http"
	"time"
)

// GET /healthcheck
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := JSONPayload{
		Message: fmt.Sprintf("healthcheck %s", time.Now()),
		Data: map[string]any{
			"status":      "available",
			"environment": app.config.env,
			"version":     app.version,
			"char":        'a',
			"b64":         []byte("This is a string"),
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
