package main

import (
	"fmt"
	"net/http"
)

func simpleHTTPServerExample() {

	/*
		Declare handlers
	*/
	http.Handle("/message", StringHandler{"Hello World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	// Equivalent without pre-declared handler
	// http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "Hello World")
	// })

	/*
		Bootstrap
	*/
	fmt.Println("Starting HTTP server on :5000")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}
}
