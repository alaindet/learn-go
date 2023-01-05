package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	version string
	config  *config
	logger  *log.Logger
}

func NewApplication(cfg *config) *application {
	return &application{
		version: version,
		config:  cfg,
		logger:  log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

func (app *application) Start(server *http.Server) {
	app.logger.Printf("starting %s server on %s", app.config.env, server.Addr)
	err := server.ListenAndServe()
	app.logger.Fatal(err)
}
