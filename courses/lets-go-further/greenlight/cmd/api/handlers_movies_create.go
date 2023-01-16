package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateMovieDto struct {
	Title   string   `json:"title"`
	Year    int32    `json:"year"`
	Runtime int32    `json:"runtime"`
	Genres  []string `json:"genres"`
}

// POST /movies
func (app *application) moviesCreateHandler(w http.ResponseWriter, r *http.Request) {

	var input CreateMovieDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.badRequestResponse(w, r, err.Error())
		return
	}

	// TODO
	fmt.Fprintf(w, "%+v\n", input)
}
