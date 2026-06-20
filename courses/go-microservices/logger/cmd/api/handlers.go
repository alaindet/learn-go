package main

import (
	"logger/data"
	"net/http"
)

type CreateLogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *App) WriteLog(w http.ResponseWriter, r *http.Request) {
	var reqData CreateLogPayload

	if err := app.ReadJSON(w, r, &reqData); err != nil {
		app.WriteJSONError(w, err, http.StatusBadRequest)
		return
	}

	event := data.LogEntry{
		Name: reqData.Name,
		Data: reqData.Data,
	}

	if err := app.Models.LogEntry.Insert(event); err != nil {
		app.WriteJSONError(w, err)
		return
	}
}
