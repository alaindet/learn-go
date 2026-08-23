package main

import (
	"errors"
	"log/slog"
	"net/http"

	"app/internal/helpers/httperr"
)

type application struct {
	config  config
	logger  *slog.Logger
	httpErr *httperr.HTTPErr
}

func initApplication(
	cfg config,
	logger *slog.Logger,
	httpErr *httperr.HTTPErr,
) *application {
	return &application{
		config:  cfg,
		logger:  logger,
		httpErr: httpErr,
	}
}

var (
	errRouteNotFound         = errors.New("cannot find route")
	errRouteMethodNotAllowed = errors.New("method not allowed on route")
)

func (app *application) handleRouteNotFound(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.httpErr.NotFound(w, r, errRouteNotFound)
}

func (app *application) handleRouteMethodNotAllowed(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.httpErr.NotFound(w, r, errRouteMethodNotAllowed)
}
