package main

import (
	commonJSON "common/json"
	"net/http"
)

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	app.WriteJSON(w, http.StatusOK, commonJSON.Response{
		Message: "Hit the broker",
	})
}
