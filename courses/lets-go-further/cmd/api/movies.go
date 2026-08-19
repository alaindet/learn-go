package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func handleCreateMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Create movie")
	})
}

func handleGetMovie() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())

		idParam := params.ByName("id") // It's an empty string if param is missing
		id, err := strconv.Atoi(idParam)
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "show details for the movie: %d\n", id)
	})
}
