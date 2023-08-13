package handlers

import (
	"errors"
	"net/http"

	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	data "greenlight/internal/data/common"
)

// GET /movies/:id
func ReadHandler(app *core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read input
		id, err := commonHttp.ReadIDRouteParameter(r)

		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Fetch
		movie, err := app.Models.Movies.Get(id)

		if errors.Is(err, data.ErrRecordNotFound) {
			app.NotFoundResponse(w, r)
			return
		}

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
			return
		}

		// Output
		data := commonHttp.JSONPayload{Data: movie}
		err = commonHttp.WriteJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
		}
	}
}
