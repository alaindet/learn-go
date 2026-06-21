package main

import (
	"bytes"
	commonJSON "common/json"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type AuthenticatePayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const logUrl = "http://logger/log"

var ErrInvalidCredentials = errors.New("Invalid credentials")

func (app *App) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload AuthenticatePayload

	if err := app.ReadJSON(w, r, &requestPayload); err != nil {
		app.WriteJSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.WriteJSONError(w, ErrInvalidCredentials, http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.WriteJSONError(w, ErrInvalidCredentials, http.StatusBadRequest)
		return
	}

	logMessage := fmt.Sprintf("%s logged in", user.Email)
	if err := app.logRequest("authentication", logMessage); err != nil {
		app.WriteJSONError(w, err)
		return
	}

	responsePayload := commonJSON.Response{
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.WriteJSON(w, http.StatusAccepted, responsePayload)
}

func (app *App) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	req, err := http.NewRequest("POST", logUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}

	if _, err := client.Do(req); err != nil {
		return err
	}

	return nil
}
