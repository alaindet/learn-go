package main

import (
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

var allowedSort = []string{
	"id",
	"-id",
	"title",
	"-title",
	"year",
	"-year",
	"runtime",
	"-runtime",
}

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
	input.Filters.SortSafelist = append(input.Filters.SortSafelist, allowedSort...)

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	movies, err := app.models.Movies.GetAll(
		input.Title,
		input.Genres,
		input.Filters,
	)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	data := JSONPayload{Data: movies}
	err = app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
