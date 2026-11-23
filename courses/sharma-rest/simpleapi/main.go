package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Config
	// TODO: Accept CLI > env > code chain
	port := 3000
	address := fmt.Sprintf(":%d", port)
	cert := "cert.pem"
	key := "key.pem"

	// Routes
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling users")
	})

	// Server config for HTTP/1.1 and HTTP/2
	serverProtocols := new(http.Protocols)
	serverProtocols.SetHTTP1(true)
	serverProtocols.SetHTTP2(true)

	server := &http.Server{
		Addr:    address,
		Handler: nil,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12, // Acceptable
		},
		Protocols:    serverProtocols,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Starting server on %s\n", address)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("could not start server", err)
		panic(err)
	}
}
