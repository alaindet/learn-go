package main

import (
	"fmt"
	"net/http"
	"time"

	"app/internal/data"
	"app/internal/helpers"
)

func handleCreateMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Create movie")
	})
}

func handleGetMovie(app *application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadIDParam(r)
		if err != nil {
			http.NotFound(w, r)
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

		err = helpers.WriteJSON(w, http.StatusOK, movie, nil)
		if err != nil {
			app.logger.Error(err.Error())
			errMessage := "The server encountered a problem and could not process your request"
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})
}
