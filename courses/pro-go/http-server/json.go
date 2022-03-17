package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonResponseExample() {
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Products)
	})

	http.Handle("/", http.RedirectHandler("/json", http.StatusTemporaryRedirect))

	/*
		Bootstrap server
	*/
	fmt.Println("Starting HTTP server on :5000")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}
}
