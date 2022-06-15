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

	// // Alternative using the built-in default servermux
	// http.HandleFunc("/snippet/view", snippetView)
	// http.HandleFunc("/snippet/create", snippetCreate)
	// http.HandleFunc("/", home)

	// Bootstrap
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
