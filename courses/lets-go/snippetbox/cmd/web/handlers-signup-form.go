package main

import (
	"net/http"

	"snippetbox.dev/internal/validator"
)

type userSignUpForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) signUpForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignUpForm{}
	data.Breadcrumbs = []*BreadcrumbLink{
		{"/", "Home", false},
		{"/users/signup", "Sign Up", true},
	}
	app.render(w, http.StatusOK, "users-signup.html", data)
}
