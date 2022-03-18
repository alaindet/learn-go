package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func inspectRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTP Request logger\n\n")

	// Method and URL
	fmt.Fprintf(w, "Method: %v\n", r.Method)
	fmt.Fprintf(w, "URL: %v\n", r.URL)

	// Headers
	for header, vals := range r.Header {
		fmt.Fprintf(w, "Header: %v: %v\n", header, vals)
	}

	fmt.Fprintln(w, strings.Repeat("=", 20))

	// Body
	data, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error reading body: %v\n", err.Error())
		return
	}

	if len(data) == 0 {
		fmt.Fprintln(w, "No body")
		return
	}

	w.Write(data)
}

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
