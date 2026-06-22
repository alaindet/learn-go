package main

import (
	"log"
	"logger/data"
	"net/http"
)

type CreateLogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *App) WriteLog(w http.ResponseWriter, r *http.Request) {

	log.Println("Attempting to write log")

	var reqData CreateLogPayload

	if err := app.ReadJSON(w, r, &reqData); err != nil {
		app.WriteJSONError(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("Log entry: name: %s, data: %v\n", reqData.Name, reqData.Data)

	event := data.LogEntry{
		Name: reqData.Name,
		Data: reqData.Data,
	}

	if err := app.Models.LogEntry.Insert(event); err != nil {
		app.WriteJSONError(w, err)
		return
	}

	log.Println("Inserted log entry")
}
