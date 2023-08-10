package main

import (
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

type MoviesListResponse struct {
	Movies   []*data.Movie `json:"movies"`
	Metadata data.Metadata `json:"metadata"`
}

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

	movies, metadata, err := app.models.Movies.GetAll(
		input.Title,
		input.Genres,
		input.Filters,
	)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	payload := JSONPayload{
		Data: MoviesListResponse{
			Movies:   movies,
			Metadata: metadata,
		},
	}

	err = app.writeJSON(w, http.StatusOK, payload, nil)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
