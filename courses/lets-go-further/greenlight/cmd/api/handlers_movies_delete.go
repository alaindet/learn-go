package main

import (
	"errors"
	"fmt"
	"net/http"

	"greenlight/internal/data"
)

// DELETE /movies/:id
// TODO?: Fetch movie first to show title in message?
func (app *application) moviesDeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Parse
	id, err := app.readIDParamHelper(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Execute
	err = app.models.Movies.Delete(id)

	// Error: 404
	if errors.Is(err, data.ErrRecordNotFound) {
		app.notFoundResponse(w, r)
		return
	}

	// Error: 500
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	// Output
	data := JSONPayload{
		Message: fmt.Sprintf("Movie #%d deleted", id),
	}
	err = app.writeJSON(w, http.StatusOK, data, nil)

	// Error: 500
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
