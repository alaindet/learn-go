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

// Implements the http.Handler interface
func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, sh.message)
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
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("URL Schema:", url.Scheme)
	fmt.Println("URL Host:", url.Host)
	fmt.Println("URL RawQuery:", url.RawQuery)
	fmt.Println("URL Path:", url.Path)
	fmt.Println("URL Fragment:", url.Fragment)
	fmt.Println("URL Hostname:", url.Hostname())
	fmt.Println("URL Port:", url.Port())
	fmt.Println("URL Query:", url.Query())
	fmt.Println("URL:", url.String())
}
