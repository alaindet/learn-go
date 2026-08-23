package httperr

import (
	"log/slog"
	"net/http"

	"app/internal/helpers"
)

type HTTPErr struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *HTTPErr {
	return &HTTPErr{logger: logger}
}

func (e HTTPErr) log(r *http.Request, err error) {
	e.logger.Error(
		err.Error(),
		"method", r.Method,
		"uri", r.URL.RequestURI(),
	)
}

func (e HTTPErr) Send(
	w http.ResponseWriter,
	r *http.Request,
	httpStatus int,
	err error,
) {
	data := helpers.JSONEnvelope{"error": err.Error()}
	writeErr := helpers.WriteJSON(w, httpStatus, data, nil)
	if writeErr != nil {
		e.log(r, writeErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (e HTTPErr) BadRequest(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	e.Send(w, r, http.StatusBadRequest, err)
}

func (e HTTPErr) NotFound(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	e.Send(w, r, http.StatusNotFound, err)
}

func (e HTTPErr) MethodNotAllowed(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	e.Send(w, r, http.StatusMethodNotAllowed, err)
}

func (e HTTPErr) InternalServerError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	e.Send(w, r, http.StatusInternalServerError, err)
}
