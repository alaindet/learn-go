package main

import (
	"database/sql"
	"fmt"
	"greenlight/internal/data"
	"net/http"
	"os"

	"log/slog"
)

type application struct {
	version string
	config  *config
	db      *sql.DB
	logger  *slog.Logger
	models  data.Models
}

func NewApplication(cfg *config) *application {

	logger := initLogger()
	db := initDabase(logger, cfg)

	return &application{
		version: version,
		config:  cfg,
		db:      db,
		logger:  logger,
		models:  data.NewModels(db),
	}
}

func initLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}

func initDabase(logger *slog.Logger, cfg *config) *sql.DB {
	db, err := openDB(cfg.db)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("database connection pool established")
	return db
}

func (app *application) Start(server *http.Server) {

	app.logger.Info(
		fmt.Sprintf("starting %s server on %s", app.config.env, server.Addr),
		"env", app.config.env,
		"address", server.Addr,
	)

	err := server.ListenAndServe()
	app.logger.Error(err.Error())
}

func (app *application) Shutdown() {
	app.db.Close()
}
