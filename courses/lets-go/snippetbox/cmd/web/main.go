package main

import (
	"log"
	"net/http"
)

func main() {
	// Create router
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Serve assets
	fileServer := http.FileServer(http.Dir("./ui/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	// Bootstrap
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
