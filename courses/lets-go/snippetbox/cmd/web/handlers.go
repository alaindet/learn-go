package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {

	// Catched a non-existing path
	// 404 Not Found
	if r.URL.Path != "/" {
		errorNotFound(w, r)
		return
	}

	templates := []string{
		"./ui/html/base.gohtml",
		"./ui/html/partials/nav.gohtml",
		"./ui/html/pages/home.gohtml",
	}

	t, err := template.ParseFiles(templates...)

	if err != nil {
		log.Println(err.Error())
		errorInternalServerError(w, r)
		return
	}

	var templateData any = nil // TODO
	err = t.ExecuteTemplate(w, "base", templateData)

	if err != nil {
		log.Println(err.Error())
		errorInternalServerError(w, r)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		errorNotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorNotAllowed(w, r, []string{http.MethodPost})
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func errorInternalServerError(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("%d Internal Server Error", http.StatusInternalServerError)
	http.Error(w, msg, http.StatusInternalServerError)
}

func errorNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func errorNotAllowed(w http.ResponseWriter, r *http.Request, methods []string) {
	w.Header().Set("Allow", strings.Join(methods, ", ")) // Set any header before writing!
	msg := fmt.Sprintf("%d Method Not Allowed", http.StatusMethodNotAllowed)
	http.Error(w, msg, http.StatusMethodNotAllowed)
}
