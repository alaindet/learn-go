package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type exceptionStep struct {
	step
}

type executer interface {
	execute() (string, error)
}

func newExceptionStep(
	name, exe, message, projectDir string,
	args []string,
) exceptionStep {
	s := exceptionStep{}
	s.step = newStep(name, exe, message, projectDir, args)
	return s
}

func (s exceptionStep) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Dir = s.projectDir
	err := cmd.Run()

	if err != nil {
		return "", newStepErr(s.name, "failed to execute", err)
	}

	if out.Len() > 0 {
		msg := fmt.Sprintf("invalid format: %s", out.String())
		return "", newStepErr(s.name, msg, nil)
	}

	return s.message, nil
}
