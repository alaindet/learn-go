package main

import (
	"fmt"
	"net/http"
	"time"
)

func NewServer(routes http.Handler, config *config) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
