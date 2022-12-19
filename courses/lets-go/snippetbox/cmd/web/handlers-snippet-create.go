package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

// TODO: Show page with HTML form
func (app *application) snippetCreateForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Breadcrumbs = []*BreadcrumbLink{
		{"/", "Home", false},
		{"/snippets/new", "Create new snippet", true},
	}
	app.render(w, http.StatusOK, "create.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Limit the POST body at 4Kb, default is 10 Mb usually
	r.Body = http.MaxBytesReader(w, r.Body, 4096)

	// Parse input
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Validation
	fieldErrors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field is required"
	}

	if utf8.RuneCountInString(title) > 100 {
		fieldErrors["title"] = "This field must be less than 100 characters long"
	}

	if strings.TrimSpace(content) == "" {
		fieldErrors["content"] = "This field is required"
	}

	if expires != 1 && expires != 7 && expires != 365 {
		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}

	// TODO: Show validation errors in a friendly way
	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

	// Save into database
	snippetId, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// https://en.wikipedia.org/wiki/Post/Redirect/Get
	http.Redirect(w, r, fmt.Sprintf("/snippets/view/%d", snippetId), http.StatusSeeOther)
}
