package main

import (
	"fmt"
	"net/http"

	"greenlight/internal/data"
	"greenlight/internal/validator"
)

// POST /movies
func (app *application) moviesCreateHandler(w http.ResponseWriter, r *http.Request) {

	var input data.CreateMovieDTO
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	data.ValidateCreateMovieDTO(v, input)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// TODO
	fmt.Fprintf(w, "%+v\n", input)
}
