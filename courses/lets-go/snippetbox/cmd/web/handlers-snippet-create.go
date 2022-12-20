package main

import (
	"fmt"
	"net/http"

	"snippetbox.dev/internal/validator"
)

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Limit the POST body at 8Kb, default is 10 Mb usually
	r.Body = http.MaxBytesReader(w, r.Body, 8192)

	var form snippetCreateForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Validation
	// TODO: Simplify or move validation to middleware
	form.Check(
		"title",
		validator.Required(form.Title),
		"This field is required",
	)

	form.Check(
		"title",
		validator.MaxChars(form.Title, 100),
		"This field must be less than 100 characters long",
	)

	form.Check(
		"content",
		validator.Required(form.Content),
		"This field is required",
	)

	form.Check(
		"expires-in-days",
		validator.InInts(form.ExpiresInDays, 1, 7, 365),
		"This field must equal 1, 7 or 365",
	)

	// Render form again, with validation errors
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		data.Breadcrumbs = []*BreadcrumbLink{
			{"/", "Home", false},
			{"/snippets/new", "Create new snippet", true},
		}
		app.render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	// Save into database
	snippetId, err := app.snippets.Insert(form.Title, form.Content, form.ExpiresInDays)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// https://en.wikipedia.org/wiki/Post/Redirect/Get
	http.Redirect(w, r, fmt.Sprintf("/snippets/view/%d", snippetId), http.StatusSeeOther)
}
