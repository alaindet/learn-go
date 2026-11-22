package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = "3000"

func main() {
	addr := "127.0.0.1:" + port

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server")
	})

	fmt.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln("error starting server", err)
	}
}
