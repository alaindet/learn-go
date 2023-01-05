package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// TODO: Automate this
const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	version string
	config  config
	logger  *log.Logger
}

func main() {

	// Config via CLI
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// App initialization
	app := &application{
		version: version,
		config:  cfg,
		logger:  log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/v1.0/healthcheck", app.healthcheckHandler)

	// Server initialization
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start server
	app.logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	app.logger.Fatal(err)
}
