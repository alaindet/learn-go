package handlers

import (
	"net/http"

	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	data "greenlight/internal/data/common"
	"greenlight/internal/data/movies"
	"greenlight/internal/validator"
)

type MoviesListResponse struct {
	Movies   []*movies.Movie         `json:"movies"`
	Metadata data.PaginationMetadata `json:"metadata"`
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
func ListHandler(app *core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var input movies.ListMoviesData

		v := validator.New()
		qs := r.URL.Query()

		input.Title = commonHttp.ReadStringFromQueryString(qs, "title", "")
		input.Genres = commonHttp.ReadCSVFromQueryString(qs, "genres", []string{})
		input.Filters.Page = commonHttp.ReadIntFromQueryString(qs, "page", 1, v)
		input.Filters.PageSize = commonHttp.ReadIntFromQueryString(qs, "page_size", 20, v)
		input.Filters.Sort = commonHttp.ReadStringFromQueryString(qs, "sort", "id")
		input.Filters.SortSafelist = append(input.Filters.SortSafelist, allowedSort...)

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			app.FailedValidationResponse(w, r, v.Errors)
			return
		}

		movies, metadata, err := app.Models.Movies.GetAll(
			input.Title,
			input.Genres,
			input.Filters,
		)

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		payload := commonHttp.JSONPayload{
			Data: MoviesListResponse{
				Movies:   movies,
				Metadata: metadata,
			},
		}

		err = commonHttp.WriteJSON(w, http.StatusOK, payload, nil)

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
		}
	}
}
