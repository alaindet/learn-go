package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	})
}
