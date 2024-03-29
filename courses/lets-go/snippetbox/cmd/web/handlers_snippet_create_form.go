package main

import (
	"net/http"

	"snippetbox.dev/internal/validator"
)

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	ExpiresInDays       int    `form:"expires-in-days"`
	validator.Validator `form:"-"`
}

func (app *application) snippetCreateForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = snippetCreateForm{
		ExpiresInDays: 365,
	}
	app.render(w, http.StatusOK, "snippet-create.html", data)
}
