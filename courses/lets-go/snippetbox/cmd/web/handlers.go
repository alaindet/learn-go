package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {

	// Catched a non-existing path
	// 404 Not Found
	if r.URL.Path != "/" {
		notFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		notFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		notAllowed(w, r, []string{http.MethodPost})
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func notAllowed(w http.ResponseWriter, r *http.Request, methods []string) {

	// Set any header before writing!
	w.Header().Set("Allow", strings.Join(methods, ", "))

	msg := fmt.Sprintf("%d Method Not Allowed", http.StatusMethodNotAllowed)
	http.Error(w, msg, http.StatusMethodNotAllowed)
}
