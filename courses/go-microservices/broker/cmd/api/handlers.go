package main

import (
	commonJSON "common/json"
	"errors"
	"net/http"
)

var (
	ErrUnknownAction      = errors.New("unknown action")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAuthService        = errors.New("failed to authenticate")
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth"`
	Log    LogPayload  `json:"log"`
	Mail   MailPayload `json:"mail"`
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
	case "log":
		app.logItem(w, requestPayload.Log)
	case "mail":
		app.sendMail(w, requestPayload.Mail)
	default:
		app.WriteJSONError(w, ErrUnknownAction, http.StatusBadRequest)
	}
}
