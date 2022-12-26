package main

import (
	"net/http"

	"snippetbox.dev/internal/validator"
)

type userSignInForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) signInForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignUpForm{}
	data.Breadcrumbs = []*BreadcrumbLink{
		{"/", "Home", false},
		{"/users/signin", "Sign In", true},
	}
	app.render(w, http.StatusOK, "users-signin.html", data)
}
