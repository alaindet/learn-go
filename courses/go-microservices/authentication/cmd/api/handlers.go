package main

import (
	"common/json"
	"errors"
	"fmt"
	"net/http"
)

type AuthenticatePayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.WriteJSONError(w, ErrInvalidCredentials, http.StatusBadRequest)
	}

	responsePayload := json.Response{
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.WriteJSON(w, http.StatusAccepted, responsePayload)
}
