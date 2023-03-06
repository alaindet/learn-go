package main

import (
	"fmt"
	"os/exec"
)

func run(cfg config) error {
	if cfg.projectDir == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	// Building multiple packages does not produce an output file (???)
	// Here, the given project (so the "main" package) and the "errors" package
	// Are built together
	args := []string{"build", ".", "errors"}
	cmd := exec.Command("go", args...)
	cmd.Dir = cfg.projectDir

	err := cmd.Run()
	if err != nil {
		return newStepErr("go build", "go build failed", err)
	}

	_, err = fmt.Fprintln(cfg.out, "Go Build: SUCCESS")
	return err
}
