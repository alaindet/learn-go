package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 3000
	address := fmt.Sprintf(":%d", port)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling users")
	})

	fmt.Printf("Starting server on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
