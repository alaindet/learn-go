package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP Server")
	http.ListenAndServe(":5000", nil)
}
