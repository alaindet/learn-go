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

	// Bootstrap
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}
