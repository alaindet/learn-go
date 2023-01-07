package main

import (
	"net/http"
	"time"

	"greenlight/internal/data"
)

func (app *application) moviesShowHandler(w http.ResponseWriter, r *http.Request) {

	// Read input
	id, err := app.readIDParamHelper(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Mock
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
