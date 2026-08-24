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
		Title   string       `json:"title"`
		Year    int          `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := helpers.ReadJSON[requestData](w, r)
		if err != nil {
			app.httpErr.BadRequest(w, r, err)
			return
		}

		movie := data.Movie{
			Title:   input.Title,
			Year:    input.Year,
			Runtime: input.Runtime,
			Genres:  input.Genres,
		}

		if errs := data.ValidateMovie(movie); len(errs) > 0 {
			app.httpErr.FailedValidation(w, r, errs)
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
