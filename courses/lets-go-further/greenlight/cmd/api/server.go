package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	serverReadTimeout     = 10 * time.Second
	serverWriteTimeout    = 20 * time.Second
	serverShutdownTimeout = 20 * time.Second
)

func (app *application) StartNewServer() error {

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ErrorLog:     log.New(os.Stdout, "", log.Ldate|log.Ltime), // TODO: Use slog
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
	}

	shutdownErr := make(chan error)

	// Background goroutine for catching signals
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit // Acts as a block until a signal is received
		app.logger.Info("shutting down server", "signal", s.String())
		ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
		defer cancel()
		shutdownErr <- server.Shutdown(ctx)
	}()

	app.logger.Info(
		fmt.Sprintf("starting %s server on %s", app.config.env, server.Addr),
		"env", app.config.env,
		"address", server.Addr,
	)

	// Skip known initial http.ErrServerClosed false error when starting the server
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErr
	if err != nil {
		return err
	}

	app.logger.Info("server stopped", "addr", server.Addr)

	return nil
}
