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

	// Validate
	v := validator.New()
	input.Validate(v)

	// Error: 422
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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

	// Execute

	// TODO
	movie.Title = input.Title
	movie.Year = input.Year
	movie.Runtime = input.Runtime
	movie.Genres = input.Genres

	err = app.models.Movies.Update(movie)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
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
