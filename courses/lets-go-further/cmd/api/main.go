package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"app/internal/helpers/httperr"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

func main() {
	// Input
	config := readConfig()

	// Initialize app
	logger := initLogger()
	httpErr := httperr.New(logger)
	app := initApplication(config, logger, httpErr)
	server := initServer(app)

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

func initServer(app *application) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}
}
