package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation failed")
)

type stepErr struct {
	step  string
	msg   string
	cause error
}

func newStepErr(step, msg string, err error) *stepErr {
	return &stepErr{step: step, msg: msg, cause: err}
}

func (s *stepErr) Error() string {
	return fmt.Sprintf("[Step %q] %s > CAUSE: %v", s.step, s.msg, s.cause)
}

func (s *stepErr) Is(target error) bool {

	// Try to assert it as a stepErr
	t, ok := target.(*stepErr)

	// You provided an error, but it's not a stepErr
	if !ok {
		return false
	}

	return t.step == s.step
}

// Used by errors.Is()
func (s *stepErr) Unwrap() error {
	return s.cause
}
