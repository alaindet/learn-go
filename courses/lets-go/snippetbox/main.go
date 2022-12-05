package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}
func main() {
	// Create and configure a router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	port := 4000
	addr := fmt.Sprintf(":%d", port)

	// Start server
	log.Printf("Starting server on :%d", port)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
