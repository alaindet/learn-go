package main

import (
	"net/http"

	"snippetbox.dev/internal/validator"
)

type snippetCreateForm struct {
	Title   string
	Content string
	Expires int
	validator.Validator
}

func (app *application) snippetCreateForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = snippetCreateForm{
		Expires: 365,
	}
	data.Breadcrumbs = []*BreadcrumbLink{
		{"/", "Home", false},
		{"/snippets/new", "Create new snippet", true},
	}
	app.render(w, http.StatusOK, "create.html", data)
}
