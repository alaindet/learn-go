package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type StringHandler struct {
	message string
}

func inspectHTTPRequest(request *http.Request) {
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("Method:", request.Method)
	fmt.Println("URL:", request.URL)
	fmt.Println("HTTP Version:", request.Proto)
	fmt.Println("Host:", request.Host)

	for headerName, headerValue := range request.Header {
		fmt.Printf("Header - Name: %s, Value: %v\n", headerName, headerValue)
	}
}

func inspectHTTPRequestURL(url *url.URL) {
	// TODO...
	fmt.Println("URL Inspection...")
}

// Implements the http.Handler interface
func (sh StringHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// inspectHTTPRequest(request)
	inspectHTTPRequestURL(request.URL)

	io.WriteString(writer, sh.message)
}

func main() {
	staticMessage := "Hello World!"
	err := http.ListenAndServe(":5000", StringHandler{message: staticMessage})

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
