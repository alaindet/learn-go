package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

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
