package main

import (
	"fmt"
	"net/http"
	"time"
)

func NewServer(mux *http.ServeMux, config *config) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
