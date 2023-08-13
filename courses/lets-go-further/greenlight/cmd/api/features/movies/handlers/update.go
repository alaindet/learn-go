package handlers

import (
	"errors"
	"fmt"
	"net/http"

	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	data "greenlight/internal/data/common"
	"greenlight/internal/data/movies"
	"greenlight/internal/validator"
)

// PATCH /movies/:id
func UpdateHandler(app *core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse
		id, err := commonHttp.ReadIDRouteParameter(r)
		if err != nil {
			app.NotFoundResponse(w, r)
			return
		}

		var input movies.UpdateMovieData
		err = commonHttp.ReadJSON(w, r, &input)

		// Error: 400
		if err != nil {
			app.BadRequestResponse(w, r, err)
			return
		}

		// Fetch
		movie, err := app.Models.Movies.Get(id)

		// Error: 404
		if errors.Is(err, data.ErrRecordNotFound) {
			app.NotFoundResponse(w, r)
			return
		}

		// Error: 500
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
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
		if movies.ValidateMovie(v, movie); !v.Valid() {
			app.FailedValidationResponse(w, r, v.Errors)
			return
		}

		// Write on database
		// See /docs/api/err-edit-conflict.sh example
		err = app.Models.Movies.Update(movie)

		if errors.Is(err, data.ErrEditConflict) {
			app.EditConflictResponse(w, r)
			return
		}

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		// Output
		data := commonHttp.JSONPayload{
			Message: fmt.Sprintf("Movie #%d %q updated", movie.ID, movie.Title),
			Data:    movie,
		}
		err = commonHttp.WriteJSON(w, http.StatusOK, data, nil)

		// Error: 500
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
		}
	}
}
