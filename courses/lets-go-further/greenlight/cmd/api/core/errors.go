package core

import (
	"fmt"
	"net/http"

	commonHttp "greenlight/cmd/api/common/http"
)

func (app *Application) LogError(r *http.Request, err error) {
	app.Logger.Error(
		err.Error(),
		"method", r.Method,
		"url", r.URL.String(),
	)
}

func (app *Application) ErrResponse(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
	errors map[string]string,
) {
	payload := commonHttp.JSONPayload{
		Message: message,
		Data:    nil,
		Error:   errors,
	}

	err := commonHttp.WriteJSON(w, status, payload, nil)

	if err != nil {
		app.LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) InternalServerErrorResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.LogError(r, err)
	status := http.StatusInternalServerError
	message := "the server encountered a problem and could not process your request"
	app.ErrResponse(w, r, status, message, nil)
}

func (app *Application) NotFoundResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	status := http.StatusNotFound
	message := "the requested resource could not be found"
	app.ErrResponse(w, r, status, message, nil)
}

func (app *Application) MethodNotAllowedResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	status := http.StatusMethodNotAllowed
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.ErrResponse(w, r, status, message, nil)
}

func (app *Application) BadRequestResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	status := http.StatusBadRequest
	message := err.Error()
	app.ErrResponse(w, r, status, message, nil)
}

func (app *Application) FailedValidationResponse(
	w http.ResponseWriter,
	r *http.Request,
	errors map[string]string,
) {
	status := http.StatusUnprocessableEntity
	message := "validation failed"
	app.ErrResponse(w, r, status, message, errors)
}

func (app *Application) EditConflictResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	status := http.StatusConflict
	message := "unable to update the record due to an edit conflict, please try again"
	app.ErrResponse(w, r, status, message, nil)
}

func (app *Application) RateLimitExceededResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	status := http.StatusTooManyRequests
	message := "rate limit exceeded"
	app.ErrResponse(w, r, status, message, nil)
}
