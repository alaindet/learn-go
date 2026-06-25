package main

import (
	"bytes"
	commonJSON "common/json"
	"encoding/json"
	"net/http"
)

const authenticateUrl = "http://authentication/authenticate"

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *App) authenticate(w http.ResponseWriter, p AuthPayload) {
	// Convert auth payload to JSON
	jsonReq, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}

	// Build a direct HTTP request
	req, err := http.NewRequest("POST", authenticateUrl, bytes.NewBuffer(jsonReq))
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}

	// Reach the authentication service
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}
	defer res.Body.Close()

	// Any other HTTP error goes here
	if res.StatusCode != http.StatusAccepted {

		// Specific 401 Unauthorized error
		if res.StatusCode == http.StatusUnauthorized {
			app.WriteJSONError(w, ErrInvalidCredentials)
			return
		}

		// Non-401 Unauthorized HTTP error
		app.WriteJSONError(w, ErrAuthService)
		return
	}

	var jsonAuthRes commonJSON.Response
	err = json.NewDecoder(res.Body).Decode(&jsonAuthRes)
	if err != nil {
		app.WriteJSONError(w, err)
		return
	}

	if jsonAuthRes.Error {
		app.WriteJSONError(w, err, http.StatusUnauthorized)
	}

	var jsonRes commonJSON.Response
	jsonRes.Message = "Authenticated!"
	jsonRes.Data = jsonAuthRes.Data

	app.WriteJSON(w, http.StatusAccepted, jsonRes)
}
