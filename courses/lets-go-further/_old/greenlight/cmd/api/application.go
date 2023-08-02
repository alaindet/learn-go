package main

import (
	"database/sql"
	"greenlight/internal/data"
	"log"
	"net/http"
	"os"
)

type application struct {
	version string
	config  *config
	db      *sql.DB
	logger  *log.Logger
	models  data.Models
}

func NewApplication(cfg *config) *application {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg.db)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("database connection pool established")

	return &application{
		version: version,
		config:  cfg,
		db:      db,
		logger:  logger,
		models:  data.NewModels(db),
	}
}

func (app *application) Start(server *http.Server) {
	app.logger.Printf("starting %s server on %s", app.config.env, server.Addr)
	err := server.ListenAndServe()
	app.logger.Fatal(err)
}

func (app *application) Shutdown() {
	app.db.Close()
}
