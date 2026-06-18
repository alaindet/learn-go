package main

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if app.ReadJSON(w, http.StatusOK, requestPayload); err != nil {
		app.WriteJSONErr(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		pp.WriteJSONErr(w, ErrInvalidCredentials, http.StatusBadRequest)
	}
}
