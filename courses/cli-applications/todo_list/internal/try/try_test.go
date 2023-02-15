package try_test

import (
	"errors"
	"testing"

	"todo_list/internal/try"
)

func TestTry(t *testing.T) {
	expected := errors.New("Break")

	err := try.Try(func(try try.Catcher) {
		try(ok("This works"))
		try(ko("Break"))
		try(ok("This never runs"))
	})

	if err == nil || err.Error() != expected.Error() {
		t.Errorf("Expected %q, got %q instead", expected, err)
	}
}

func ok(message string) error {
	return nil
}

func ko(message string) error {
	return errors.New(message)
}
