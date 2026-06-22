package main

import (
	"bytes"
	commonJSON "common/json"
	"encoding/json"
	"net/http"
)

const logUrl = "http://logger/log"

// TODO: Not working
func (app *App) logItem(w http.ResponseWriter, p LogPayload) {
	// Convert auth payload to JSON
	jsonReq, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}

	// Build a direct HTTP request
	req, err := http.NewRequest("POST", logUrl, bytes.NewBuffer(jsonReq))
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}
	defer res.Body.Close()

	// Any other HTTP error goes here
	if res.StatusCode != http.StatusAccepted {
		// Non-401 Unauthorized HTTP error
		app.WriteJSONError(w, ErrAuthService)
		return
	}

	var resData commonJSON.Response
	resData.Message = "logged"

	app.WriteJSON(w, http.StatusAccepted, resData)
}
