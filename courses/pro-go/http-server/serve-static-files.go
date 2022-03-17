package main

import (
	"fmt"
	"net/http"
)

func serveStaticFilesExample() {
	/*
		Serve static files
	*/
	fsHandler := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fsHandler))

	/*
		Declare HTTP handlers
	*/
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/assets/templates/hello.html", http.StatusTemporaryRedirect))

	/*
		Bootstrap server
	*/
	fmt.Println("Starting HTTP server on :5000")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}
}
