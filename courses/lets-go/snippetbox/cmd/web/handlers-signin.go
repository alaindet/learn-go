package main

import (
	"errors"
	"net/http"

	"snippetbox.dev/internal/models"
	"snippetbox.dev/internal/validator"
)

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {

	// Parse input
	var form userSignInForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Validation
	// TODO: Simplify or move validation to middleware
	form.Check(
		"email",
		validator.Required(form.Email),
		"This field is required",
	)

	form.Check(
		"email",
		validator.IsEmail(form.Email),
		"This field must be a valid email",
	)

	form.Check(
		"password",
		validator.Required(form.Password),
		"This field is required",
	)

	// Render form again, with validation errors
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		data.Breadcrumbs = []*BreadcrumbLink{
			{"/", "Home", false},
			{"/users/signin", "Sign In", true},
		}
		app.render(w, http.StatusUnprocessableEntity, "users-signin.html", data)
		return
	}

	userId, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("Email/password incorrect")
			data := app.newTemplateData(r)
			data.Form = form
			data.Breadcrumbs = []*BreadcrumbLink{
				{"/", "Home", false},
				{"/users/signin", "Sign In", true},
			}
			app.render(w, http.StatusUnprocessableEntity, "users-signin.html", data)
		} else {
			app.serverError(w, err)
		}
	}

	// It's best to refresh the token whenever the user changes privileges
	// https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Session_Management_Cheat_Sheet.md#renew-the-session-id-after-any-privilege-level-change
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), sessionKeyFlash, "You successfully signed in.")
	app.sessionManager.Put(r.Context(), sessionKeyUserId, userId)

	// https://en.wikipedia.org/wiki/Post/Redirect/Get
	http.Redirect(w, r, "/snippets/new", http.StatusSeeOther)
}
