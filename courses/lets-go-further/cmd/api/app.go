package main

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"app/internal/helpers/httperr"
)

type application struct {
	config  config
	logger  *slog.Logger
	httpErr *httperr.HTTPErr
	db      *sql.DB
}

func initApplication(cfg config) *application {

	// Logger
	logger := initLogger()

	// HTTP Error
	httpErr := httperr.New(logger)

	// Database
	db, err := openDB(logger, cfg.database)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	db.Close()
	logger.Info("database connection pool established")

	// App
	return &application{
		config:  cfg,
		logger:  logger,
		httpErr: httpErr,
		db:      db,
	}
}

func initLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
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
