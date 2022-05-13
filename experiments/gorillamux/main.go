package main

import (
	"fmt"
	"log"
	"net/http"
)

var port string = "8080"
var db *DB

func main() {
	db = NewDB()
	router := NewRouter()
	fmt.Printf("Boostrapped server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
