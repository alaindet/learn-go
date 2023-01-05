package main

import (
	"fmt"
	"net/http"
)

func (app *application) moviesShowHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParamHelper(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id)

}
