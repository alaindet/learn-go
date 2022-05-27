package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {

	expected := "$70200.00"

	// Setup: mock stdout
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test
	main()

	// Cleanup: Restore stdout
	_ = w.Close()
	rawResult, _ := io.ReadAll(r)
	os.Stdout = stdOut
	result := string(rawResult)

	if !strings.Contains(result, expected) {
		t.Errorf("Wrong balance returned. Expected %q to be found in %q", expected, result)
	}
}
