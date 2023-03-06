package main

import "os/exec"

type step struct {
	name       string
	exe        string
	args       []string
	message    string
	projectDir string
}

func newStep(name, exe, message, projectDir string, args []string) step {
	return step{
		name:       name,
		exe:        exe,
		message:    message,
		args:       args,
		projectDir: projectDir,
	}
}

func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.projectDir

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	return s.message, nil
}
