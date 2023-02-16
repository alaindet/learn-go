package input_test

import (
	"os"
	"testing"
	"todo_list/internal/input"
)

func TestCLIFlags(t *testing.T) {

	foo := "more than a word"
	bar := 42
	baz := true

	f := input.NewCLIFlags()
	f.String("foo", "", "The value of foo")
	f.Int("bar", -1, "The value of bar")
	f.Bool("baz", false, "The value of baz")

	// Simulate STDIN
	os.Args = []string{"cmd", "-foo", "more than a word", "-bar", "42", "-baz"}

	f.Parse()

	if foo != f.GetString("foo") {
		t.Errorf("Expected %q, got %q", foo, f.GetString("foo"))
	}

	if bar != f.GetInt("bar") {
		t.Errorf("Expected %d, got %d", bar, f.GetInt("bar"))
	}

	if baz != f.GetBool("baz") {
		t.Errorf("Expected %t, got %t", baz, f.GetBool("baz"))
	}
}
