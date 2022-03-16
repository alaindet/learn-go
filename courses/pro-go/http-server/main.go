package main

import (
	"fmt"
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

// Implements the http.Handler interface
func (sh StringHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	io.WriteString(writer, sh.message)
}

func main() {
	staticMessage := "Hello World!"
	err := http.ListenAndServe(":5000", StringHandler{message: staticMessage})

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
