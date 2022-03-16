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

// Implements the http.Handler interface
func (sh StringHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// inspectHTTPRequest(request)
	// inspectHTTPRequestURL(request.URL)

	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/message":
		io.WriteString(writer, sh.message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
}

func main() {
	staticMessage := "Hello World!"
	err := http.ListenAndServe(":5000", StringHandler{message: staticMessage})

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
