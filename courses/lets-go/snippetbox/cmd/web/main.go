package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

type application struct {
	config
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *sql.DB
}

func main() {

	// Application
	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		config:   *ConfigNew(),
	}

	// Init DB connection
	db, err := app.openDB()
	if err != nil {
		app.errorLog.Fatal(err)
	}
	defer db.Close()
	app.db = db

	// Bootstrap
	webServer := &http.Server{
		Addr:     app.config.addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("starting %s on %s\n", app.config.name, app.config.addr)
	err = webServer.ListenAndServe()
	app.errorLog.Fatal(err)
}
