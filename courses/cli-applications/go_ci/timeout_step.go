package main

import (
	"context"
	"os/exec"
	"time"
)

type timeoutStep struct {
	step
	timeout time.Duration
}

func newTimeoutStep(
	name, exe, message, projectDir string,
	args []string,
	timeout time.Duration,
) timeoutStep {
	s := timeoutStep{}
	s.step = newStep(name, exe, message, projectDir, args)
	s.timeout = timeout

	// Defaults to 10 seconds
	if s.timeout == 0 {
		s.timeout = 10 * time.Second
	}

	return s
}

// Store it here, so you can swap it for test mocks
var command = exec.CommandContext

func (s timeoutStep) execute() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	cmd := command(ctx, s.exe, s.args...)
	cmd.Dir = s.projectDir
	err := cmd.Run()

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			cause := context.DeadlineExceeded
			return "", newStepErr(s.name, "failed time out", cause)
		}

		return "", newStepErr(s.name, "failed to execute", err)
	}

	return s.message, nil
}
