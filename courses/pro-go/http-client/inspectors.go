package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func inspectRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTP Request logger\n\n")

	// Method and URL
	fmt.Fprintf(w, "Method: %v\n", r.Method)
	fmt.Fprintf(w, "URL: %v\n", r.URL)

	// Headers
	for header, vals := range r.Header {
		fmt.Fprintf(w, "Header: %v: %v\n", header, vals)
	}

	fmt.Fprintln(w, strings.Repeat("=", 20))

	// Body
	data, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error reading body: %v\n", err.Error())
		return
	}

	if len(data) == 0 {
		fmt.Fprintln(w, "No body")
		return
	}

	w.Write(data)
}

func inspectResponse(r *http.Response) {
	url, err := r.Location()
	_ = err
	fmt.Printf("URL: %v\n", url)
	fmt.Printf("Status: %d %s\n", r.StatusCode, r.Status)
	fmt.Printf("Content-Length: %v\n", r.ContentLength)
	fmt.Printf("Cookies: %v\n", r.Cookies())
	fmt.Printf("Headers:\n%v\n", r.Header)
	fmt.Printf("Body:\n%v\n", r.Body)
}
