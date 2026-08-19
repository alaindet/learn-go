package main

import (
	"fmt"
	"net/http"
)

func handleCreateMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Create movie")
	})
}

func handleGetMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := readIDParam(r)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "show details for the movie: %d\n", id)
	})
}
