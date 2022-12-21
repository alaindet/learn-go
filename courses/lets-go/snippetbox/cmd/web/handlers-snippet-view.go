package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"snippetbox.dev/internal/models"
)

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
	data.Flash = app.sessionManager.PopString(r.Context(), flashKey)
	data.Breadcrumbs = []*BreadcrumbLink{
		{"/", "Home", false},
		{"/snippets/new", "Create new snippet", false},
		{"/snippets/view/" + id, "Snippet #" + id, true},
	}

	app.render(w, http.StatusOK, "snippet-view.html", data)
}
