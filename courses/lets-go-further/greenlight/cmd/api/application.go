package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"greenlight/internal/database"
)

type application struct {
	version string
	config  *config
	db      *sql.DB
	logger  *log.Logger
}

func NewApplication(cfg *config) *application {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.OpenDB(cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("database connection pool established")

	return &application{
		version: version,
		config:  cfg,
		db:      db,
		logger:  logger,
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
