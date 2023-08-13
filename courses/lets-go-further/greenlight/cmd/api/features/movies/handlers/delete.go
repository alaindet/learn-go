package handlers

import (
	"errors"
	"fmt"
	"net/http"

	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	data "greenlight/internal/data/common"
)

// DELETE /movies/:id
func DeleteHandler(app *core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse
		id, err := commonHttp.ReadIDRouteParameter(r)
		if err != nil {
			app.NotFoundResponse(w, r)
			return
		}

		// Execute
		err = app.Models.Movies.Delete(id)

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

		// Output
		data := commonHttp.JSONPayload{
			Message: fmt.Sprintf("Movie #%d deleted", id),
		}
		err = commonHttp.WriteJSON(w, http.StatusOK, data, nil)

		// Error: 500
		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
		}
	}
}
