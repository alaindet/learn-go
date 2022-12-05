package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Create and configure a router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)     // TODO: Should be GET /snippets/:id
	mux.HandleFunc("/snippet/create", snippetCreate) // TODO: Should be POST /snippets

	port := 4000
	addr := fmt.Sprintf(":%d", port)

	// Start server
	log.Printf("Starting server on :%d", port)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}
