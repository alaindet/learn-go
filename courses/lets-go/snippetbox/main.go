package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
