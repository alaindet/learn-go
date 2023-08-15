package handlers

import (
	"errors"
	"fmt"
	"net/http"

	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	"greenlight/internal/data/common"
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

		if errors.Is(err, common.ErrDuplicateEmail) {
			v.AddError("email", "email already in use")
			app.FailedValidationResponse(w, r, v.Errors)
			return
		}

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		// Send success email
		err = app.Mailer.Send(user.Email, "user_welcome.tmpl", user)
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		data := commonHttp.JSONPayload{
			Message: fmt.Sprintf("User %s successfully signed in", user.Email),
			Data:    user,
		}

		err = commonHttp.WriteJSON(w, http.StatusCreated, data, nil)

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}
	}
}
