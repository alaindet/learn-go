package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("Display a specific snippet..."))
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
	w.Header().Set("Allow", strings.Join(methods, ","))

	w.WriteHeader(405) // TODO: Status code is a header?
	w.Write([]byte("405 Method Not Allowed"))
}
