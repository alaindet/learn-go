package main

import (
	"net/http"
)

func main() {

	cfg := NewConfig()
	app := NewApplication(cfg)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/v1.0/healthcheck", app.healthcheckHandler)

	server := NewServer(mux, cfg)
	app.Start(server)
}
