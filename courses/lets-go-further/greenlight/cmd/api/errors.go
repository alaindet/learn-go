package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

func (app *application) errResponse(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
	errors map[string]string,
) {
	payload := JSONPayload{Message: message, Data: nil, Error: errors}
	err := app.writeJSON(w, status, payload, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) internalServerErrorResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errResponse(w, r, http.StatusInternalServerError, message, nil)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errResponse(w, r, http.StatusNotFound, message, nil)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errResponse(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (app *application) badRequestResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.errResponse(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (app *application) failedValidationResponse(
	w http.ResponseWriter,
	r *http.Request,
	errors map[string]string,
) {
	app.errResponse(w, r, http.StatusUnprocessableEntity, "validation failed", errors)
}

func (app *application) editConflictResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errResponse(w, r, http.StatusConflict, message, nil)
}
