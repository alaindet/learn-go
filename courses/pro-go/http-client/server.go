package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("server.go => init()")

	http.HandleFunc("/gimme-html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	http.HandleFunc("/gimme-json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Products)
	})

	http.HandleFunc("/gimme-log", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		inspectRequest(w, r)
	})
}
