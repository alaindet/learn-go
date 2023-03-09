package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func run(cfg config) error {
	if cfg.projectDir == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	pipeline := make([]executer, 4)
	errCh := make(chan error)
	doneCh := make(chan struct{})
	sigCh := make(chan os.Signal, 1)

	// Capture signals SIGINT (Signal interrupt - Ctrl + C) and
	// SIGTERM (Signal terminate) and send them to signCh channel
	// They can both be ignored by the program or used to gracefully terminate
	// SIGKILL abrutply kills a process and cannot be captured
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Build step
	pipeline[0] = newStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		cfg.projectDir,
		[]string{"build", ".", "errors"},
	)

	// Test step
	pipeline[1] = newStep(
		"go test",
		"go",
		"Go Test: SUCCESS",
		cfg.projectDir,
		[]string{"test", "-v"},
	)

	// Format step
	pipeline[2] = newExceptionStep(
		"go fmt",
		"gofmt",
		"Gofmt: SUCCESS",
		cfg.projectDir,
		[]string{"-l", "."},
	)

	// Git push
	pipeline[3] = newTimeoutStep(
		"git push",
		"git",
		"Git Push: SUCCESS",
		cfg.projectDir,
		[]string{"push", "origin", "master"},
		10*time.Second,
	)

	// Execute in a goroutine
	go func() {
		for _, s := range pipeline {
			msg, err := s.execute()
			if err != nil {
				errCh <- err
				return
			}

			_, err = fmt.Fprintln(cfg.out, msg)
			if err != nil {
				errCh <- err
				return
			}
		}
		close(doneCh)
	}()

	// Listen for errors, signals or completion
	for {
		select {
		case capturedSignal := <-sigCh:
			// Stop capturing signals
			signal.Stop(sigCh)
			return fmt.Errorf("%s: Exiting %w", capturedSignal, ErrSignal)
		case err := <-errCh:
			return err
		case <-doneCh:
			return nil
		}
	}
}
