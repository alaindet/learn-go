package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// Renders the given page into a buffer first, then writes the actual output
// If everything is OK
func (app *application) render(
	w http.ResponseWriter,
	status int,
	page string,
	templateData *templateData,
) {
	ts, ok := app.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	// Write into a buffer first
	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", templateData)

	if err != nil {
		app.serverError(w, err)
		return
	}

	// Write to real output
	w.WriteHeader(status)
	buf.WriteTo(w)
}
