package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := loadConfig()

	// Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Router
	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir(cfg.staticPath))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

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
