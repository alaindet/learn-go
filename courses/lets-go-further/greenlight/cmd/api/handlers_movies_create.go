package main

import (
	"fmt"
	"net/http"

	"greenlight/internal/data"
	"greenlight/internal/validator"
)

// POST /movies
func (app *application) moviesCreateHandler(w http.ResponseWriter, r *http.Request) {

	// Parse
	var input data.CreateMovieData
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Validate
	v := validator.New()
	input.Validate(v)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Execute
	movie := input.ToMovie()
	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	// Output
	url := fmt.Sprintf("/v1.0/movies/%d", movie.ID)
	data := JSONPayload{
		Message: fmt.Sprintf("Movie %q created with ID %d", movie.Title, movie.ID),
		Data: map[string]any{
			"url":   url,
			"movie": movie,
		},
	}
	headers := make(http.Header)
	headers.Set("Location", url)
	err = app.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
