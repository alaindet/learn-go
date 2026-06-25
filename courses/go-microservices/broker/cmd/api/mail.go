package main

import (
	"bytes"
	commonJSON "common/json"
	"encoding/json"
	"net/http"
)

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

const mailUrl = "http://mail/send"

func (app *App) sendMail(w http.ResponseWriter, p MailPayload) {
	// Convert auth payload to JSON
	jsonReq, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}

	// Build a direct HTTP request
	req, err := http.NewRequest("POST", mailUrl, bytes.NewBuffer(jsonReq))
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
	resData.Message = "email sent"

	app.WriteJSON(w, http.StatusAccepted, resData)
}
