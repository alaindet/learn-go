package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	cfg := loadConfig()

	// Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Router
	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir(cfg.staticPath))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// Bootstrap
	webServer := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("starting %s on %s\n", cfg.name, cfg.addr)
	err := webServer.ListenAndServe()
	errorLog.Fatal(err)
}
