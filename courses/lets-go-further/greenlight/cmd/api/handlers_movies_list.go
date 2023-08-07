package main

import (
	"fmt"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

// GET /movies
func (app *application) moviesListHandler(w http.ResponseWriter, r *http.Request) {

	var input data.ListMoviesData

	v := validator.New()
	qs := r.URL.Query()

	input.Title = app.readStringFromQueryString(qs, "title", "")
	input.Genres = app.readCSVFromQueryString(qs, "genres", []string{})
	input.Filters.Page = app.readIntFromQueryString(qs, "page", 1, v)
	input.Filters.PageSize = app.readIntFromQueryString(qs, "page_size", 20, v)
	input.Filters.Sort = app.readStringFromQueryString(qs, "sort", "id")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// TODO: Perform query on the database
	fmt.Fprintf(w, "%+v\n", input)
}
