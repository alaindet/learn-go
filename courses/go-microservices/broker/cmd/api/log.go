package main

import (
	"bytes"
	commonJSON "common/json"
	"encoding/json"
	"net/http"

	"broker/event"
)

const logUrl = "http://logger/log"

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

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

func (app *App) logEventViaRabbitMQ(w http.ResponseWriter, p LogPayload) {
	if err := app.pushToQueue(p.Name, p.Data); err != nil {
		app.WriteJSONError(w, err)
		return
	}

	var payload commonJSON.Response
	payload.Message = "logged via RabbitMQ"
	app.WriteJSON(w, http.StatusAccepted, payload)
}

func (app *App) pushToQueue(name, message string) error {
	emitter, err := event.NewEventEmitter(app.RabbitMQ)
	if err != nil {
		return err
	}

	payload := LogPayload{
		Name: name,
		Data: message,
	}

	jsonData, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	if err := emitter.Push(string(jsonData), event.SeverityInfo); err != nil {
		return err
	}

	return nil
}
