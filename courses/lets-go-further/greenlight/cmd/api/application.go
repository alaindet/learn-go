package main

import (
	"database/sql"
	"greenlight/internal/data"
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

func (app *application) Shutdown() {
	app.db.Close()
}
