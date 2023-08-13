package handlers

import (
	"fmt"
	"net/http"

	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	"greenlight/internal/data/movies"
	"greenlight/internal/validator"
)

// POST /movies
func CreateHandler(app *core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse
		var input movies.CreateMovieData
		err := commonHttp.ReadJSON(w, r, &input)

		// Error: 400
		if err != nil {
			app.BadRequestResponse(w, r, err)
			return
		}

		// Validate
		v := validator.New()
		input.Validate(v)

		// Error: 422
		if !v.Valid() {
			app.FailedValidationResponse(w, r, v.Errors)
			return
		}

		// Execute
		movie := input.ToMovie()
		err = app.Models.Movies.Insert(movie)

		// Error: 500
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		// Output
		url := fmt.Sprintf("/v1/movies/%d", movie.ID)
		data := commonHttp.JSONPayload{
			Message: fmt.Sprintf("Movie %q created with ID %d", movie.Title, movie.ID),
			Data: map[string]any{
				"url":   url,
				"movie": movie,
			},
		}
		headers := make(http.Header)
		headers.Set("Location", url)
		err = commonHttp.WriteJSON(w, http.StatusCreated, data, headers)

		// Error: 500
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
		}
	}
}
