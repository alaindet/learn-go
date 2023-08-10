package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func NewServer(routes http.Handler, config *config) *http.Server {

	// TODO: Use slog?
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      routes,
		IdleTimeout:  time.Minute,
		ErrorLog:     logger,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
