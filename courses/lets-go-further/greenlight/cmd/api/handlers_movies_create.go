package main

import (
	"fmt"
	"greenlight/internal/data"
	"net/http"
)

type CreateMovieDTO struct {
	Title   string       `json:"title"`
	Year    int32        `json:"year"`
	Runtime data.Runtime `json:"runtime"`
	Genres  []string     `json:"genres"`
}

// POST /movies
func (app *application) moviesCreateHandler(w http.ResponseWriter, r *http.Request) {

	var input CreateMovieDTO
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// TODO
	fmt.Fprintf(w, "%+v\n", input)
}
