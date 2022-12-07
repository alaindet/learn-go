package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := loadConfig()

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
	log.Printf("Starting Snippetbox on %s\n", cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
