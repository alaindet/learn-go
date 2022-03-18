package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func httpClientExample() {
	// Start HTTP server
	go http.ListenAndServe(":5000", nil)

	time.Sleep(time.Second * 1)

	// Perform HTTP request to running server
	response, err := http.Get("http://localhost:5000/gimme-json")

	// // Debug-only
	// inspectResponse(response)
	// response.Write(os.Stdout)

	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}

	if response.StatusCode != http.StatusOK {
		fmt.Printf("HTTP ERROR %d\n", response.StatusCode)
		return
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("ERROR: Cannot read HTTP Response body", err.Error())
	}

	defer response.Body.Close()
	os.Stdout.Write(data) // Debug-only
}
