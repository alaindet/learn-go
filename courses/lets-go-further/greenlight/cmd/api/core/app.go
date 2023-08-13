package core

import (
	"database/sql"
	"greenlight/internal/data/models"
	"os"

	"log/slog"
)

type Application struct {
	Version   string
	UrlPrefix string // Ex.: "/v1"
	Config    *Config
	Db        *sql.DB
	Logger    *slog.Logger
	Models    models.Models
}

func NewApplication(cfg *Config) *Application {

	logger := initLogger()
	db := initDabase(logger, cfg)
	models := initModels(db)
	prefix := "/" + ApiVersion

	return &Application{
		Version:   Version,
		UrlPrefix: prefix,
		Config:    cfg,
		Db:        db,
		Logger:    logger,
		Models:    models,
	}
}

func initLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}

func initDabase(logger *slog.Logger, cfg *Config) *sql.DB {
	db, err := openDB(cfg.Db)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("database connection pool established")
	return db
}

func initModels(db *sql.DB) models.Models {
	return models.NewModels(db)
}

func (app *Application) Shutdown() {
	app.Db.Close()
}
