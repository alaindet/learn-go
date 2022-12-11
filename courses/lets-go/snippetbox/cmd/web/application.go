package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"snippetbox.dev/internal/models"
)

type application struct {
	config
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *sql.DB // TODO: Remove?
	snippets *models.SnippetModel
}

func initApp() *application {

	// Init loggers
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Init config
	config := *NewConfig()

	// Init database
	db, err := openDB(config.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Application
	return &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		config:   config,
		db:       db,
		snippets: models.NewSnippetModel(db),
	}
}

func (app *application) initWebServer() *http.Server {
	return &http.Server{
		Addr:     app.config.addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}
}

func (app *application) shutdown() {
	app.db.Close()
}
