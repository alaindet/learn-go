package main

import (
	"net/http"
	"time"

	"greenlight/internal/data"
)

// GET /movies/:id
func (app *application) moviesShowHandler(w http.ResponseWriter, r *http.Request) {

	// Read input
	id, err := app.readIDParamHelper(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	payload := JSONPayload{
		Data: data.Movie{
			ID:        id,
			CreatedAt: time.Now(),
			Title:     "Casablanca",
			Runtime:   102,
			Genres:    []string{"drama", "romance", "war"},
			Version:   1,
		},
	}

	err = app.writeJSON(w, http.StatusOK, payload, nil)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
