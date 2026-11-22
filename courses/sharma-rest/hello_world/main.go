package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var personData = map[string]Person{
	"1": {Name: "John Doe", Age: 30},
	"2": {Name: "Jane Doe", Age: 28},
	"3": {Name: "Jack Doe", Age: 25},
}

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server started on port %d\n", port)
	http.HandleFunc("/person", getPersonHandler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
