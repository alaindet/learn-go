package main

import (
	"fmt"
	"net/http"
	"time"

	"app/internal/data"
	"app/internal/helpers"
)

func handleCreateMovie(app *application) http.Handler {
	type requestData struct {
		Title   string   `json:"title"`
		Year    int      `json:"year"`
		Runtime int      `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input requestData
		err := helpers.ReadJSON(w, r, &input)
		if err != nil {
			app.httpErr.BadRequest(w, r, err)
			return
		}

		// Temporary
		fmt.Fprintf(w, "%+v\n", input)
	})
}

func handleGetMovie(app *application) http.Handler {
	type responseData struct {
		Movie data.Movie `json:"movie"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadIDParam(r)
		if err != nil {
			app.httpErr.NotFound(w, r, err)
			return
		}

		movie := data.Movie{
			ID:        id,
			CreatedAt: time.Now(),
			Title:     "Casablanca",
			Runtime:   102,
			Genres:    []string{"drama", "romance", "war"},
			Version:   1,
		}

		err = helpers.WriteJSON(w, http.StatusOK, responseData{Movie: movie}, nil)
		if err != nil {
			app.httpErr.InternalServerError(w, r, err)
		}
	})
}
