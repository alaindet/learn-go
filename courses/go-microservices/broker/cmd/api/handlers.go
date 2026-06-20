package main

import (
	"bytes"
	commonJSON "common/json"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	authenticateUrl = "http://authentication/authenticate" // TODO

	ErrUnknownAction      = errors.New("unknown action")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAuthService        = errors.New("failed to authenticate")
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omiempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	app.WriteJSON(w, http.StatusOK, commonJSON.Response{
		Message: "Hit the broker",
	})
}

func (app *App) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	if err := app.ReadJSON(w, r, &requestPayload); err != nil {
		app.WriteJSONError(w, err, http.StatusBadRequest)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	default:
		app.WriteJSONError(w, ErrUnknownAction, http.StatusBadRequest)
	}
}

func (app *App) authenticate(w http.ResponseWriter, a AuthPayload) {
	// Convert auth payload to JSON
	jsonReq, err := json.MarshalIndent(a, "", "\t")
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
