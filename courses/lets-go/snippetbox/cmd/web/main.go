package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Bootstrap
	port := 4000
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting Snippetbox on %s\n", addr)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
