package main

import (
	"fmt"
)

func run(cfg config) error {
	if cfg.projectDir == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	pipeline := make([]executer, 3)

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

	for _, s := range pipeline {
		msg, err := s.execute()
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(cfg.out, msg)
		if err != nil {
			return err
		}
	}

	return nil
}
