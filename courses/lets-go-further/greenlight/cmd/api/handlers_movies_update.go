package main

import (
	"errors"
	"fmt"
	"net/http"

	"greenlight/internal/data"
	"greenlight/internal/validator"
)

// PUT /movies/:id
func (app *application) moviesUpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Parse
	id, err := app.readIDParamHelper(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input data.UpdateMovieData
	err = app.readJSON(w, r, &input)

	// Error: 400
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Fetch
	movie, err := app.models.Movies.Get(id)

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

	if input.Title != nil {
		movie.Title = *input.Title
	}

	if input.Year != nil {
		movie.Year = *input.Year
	}

	if input.Runtime != nil {
		movie.Runtime = *input.Runtime
	}

	if input.Genres != nil {
		movie.Genres = input.Genres
	}

	// Validate movie data
	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Write on database
	// See /docs/api/err-edit-conflict.sh example
	err = app.models.Movies.Update(movie)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r, err)
		default:
			app.internalServerErrorResponse(w, r, err)
		}
		return
	}

	// Output
	data := JSONPayload{
		Message: fmt.Sprintf("Movie #%d %q updated", movie.ID, movie.Title),
		Data:    movie,
	}
	err = app.writeJSON(w, http.StatusOK, data, nil)

	// Error: 500
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
