package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	// Input
	config := readConfig()

	// Initialize app
	logger := initLogger()
	app := initApplication(config, logger)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Bootstrap app
	logger.Info("starting server", "addr", server.Addr, "env", config.env)
	err := server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func readConfig() config {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
	return cfg
}

func initLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func initApplication(cfg config, logger *slog.Logger) *application {
	return &application{
		config: cfg,
		logger: logger,
	}
}
