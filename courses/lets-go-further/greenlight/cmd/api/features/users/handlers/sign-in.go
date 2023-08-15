package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"greenlight/cmd/api/common"
	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	data "greenlight/internal/data/common"
	"greenlight/internal/data/users"
	"greenlight/internal/validator"
)

// POST /users
func SignInHandler(app *core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse input JSON
		var input users.SignInUserData
		err := commonHttp.ReadJSON(w, r, &input)

		// Error: Bad Request
		if err != nil {
			app.BadRequestResponse(w, r, err)
			return
		}

		// Create a new user entity
		user, err := input.ToUser()
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		// Validate the new user
		v := validator.New()
		users.ValidateUser(v, user)

		// Error: Unprocessable Content
		if !v.Valid() {
			app.FailedValidationResponse(w, r, v.Errors)
			return
		}

		// Create on database
		err = app.Models.Users.Insert(user)

		if errors.Is(err, data.ErrDuplicateEmail) {
			v.AddError("email", "email already in use")
			app.FailedValidationResponse(w, r, v.Errors)
			return
		}

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		// Send success email in background
		common.BackgroundTask(app, func() {
			err = app.Mailer.Send(user.Email, "user_welcome.tmpl", user)
			if err != nil {
				app.Logger.Error(err.Error(), nil)
			}
		})

		data := commonHttp.JSONPayload{
			Message: fmt.Sprintf("User %s successfully signed in", user.Email),
			Data:    user,
		}

		err = commonHttp.WriteJSON(w, http.StatusAccepted, data, nil)

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}
	}
}
