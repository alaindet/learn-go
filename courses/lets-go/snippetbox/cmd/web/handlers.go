package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"snippetbox.dev/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets
	data.Breadcrumbs = []*BreadcrumbLink{
		{"/", "Home", true},
		{"/snippets/new", "Create new snippet", false},
	}

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	// TODO: Read the docs
	params := httprouter.ParamsFromContext(r.Context())

	id := params.ByName("id")

	if id == "" {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet
	data.Breadcrumbs = []*BreadcrumbLink{
		{"/", "Home", false},
		{"/snippets/new", "Create new snippet", false},
		{"/snippets/view/" + id, "Snippet #" + id, true},
	}

	app.render(w, http.StatusOK, "snippet-view.html", data)
}

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

	// Validation
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
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
