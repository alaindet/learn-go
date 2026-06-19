package main

import (
	"bytes"
	commonJSON "common/json"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrUnknownAction = errors.New("Unknown action")
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

	if err := app.ReadJSON(w, r, requestPayload); err != nil {
		app.WriteJSONError(w, err, http.StatusBadRequest)
		return
	}

	switch requestPayload.Action {
	case "auth":
		// TODO...
	default:
		app.WriteJSONError(w, ErrUnknownAction, http.StatusBadRequest)
	}
}

func (app *App) authenticate(w http.ResponseWriter, a AuthPayload) error {
	jsonData, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	return nil
}
