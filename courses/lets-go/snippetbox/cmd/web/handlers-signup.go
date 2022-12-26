package main

import (
	"errors"
	"net/http"

	"snippetbox.dev/internal/models"
	"snippetbox.dev/internal/validator"
)

func (app *application) signUp(w http.ResponseWriter, r *http.Request) {

	// Parse input
	var form userSignUpForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Validation
	// TODO: Simplify or move validation to middleware
	form.Check(
		"name",
		validator.Required(form.Name),
		"This field is required",
	)

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

	form.Check(
		"password",
		validator.MinChars(form.Password, 8),
		"This field must be at least 8 characters long",
	)

	// Render form again, with validation errors
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		data.Breadcrumbs = []*BreadcrumbLink{
			{"/", "Home", false},
			{"/users/signup", "Sign Up", true},
		}
		app.render(w, http.StatusUnprocessableEntity, "users-signup.html", data)
		return
	}

	// Save into database
	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {

		// Email already exists?
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already in use")
			data := app.newTemplateData(r)
			data.Form = form
			data.Breadcrumbs = []*BreadcrumbLink{
				{"/", "Home", false},
				{"/users/signup", "Sign Up", true},
			}
			app.render(w, http.StatusUnprocessableEntity, "users-signup.html", data)
		} else {
			app.serverError(w, err)
		}
	}

	// Notify the user
	app.sessionManager.Put(
		r.Context(),
		sessionKeyFlash,
		"You successfully signed up on Snippetbox. Welcome!",
	)

	// https://en.wikipedia.org/wiki/Post/Redirect/Get
	http.Redirect(w, r, "/users/signin", http.StatusSeeOther)
}
