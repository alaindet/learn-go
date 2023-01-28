package main

import (
	"errors"
	"net/http"

	"greenlight/internal/data"
)

// GET /movies/:id
func (app *application) moviesShowHandler(w http.ResponseWriter, r *http.Request) {

	// Read input
	id, err := app.readIDParamHelper(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Fetch
	movie, err := app.models.Movies.Get(id)

	if errors.Is(err, data.ErrRecordNotFound) {
		app.notFoundResponse(w, r)
		return
	}

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	// Output
	data := JSONPayload{Data: movie}
	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
