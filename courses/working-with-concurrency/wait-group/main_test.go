package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printIt(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"alpha", "alpha"},
		{"beta", "beta"},
		{"gamma", "gamma"},
		{"omega", "omega"},
	}

	// Setup: mock stdout
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test
	var wg sync.WaitGroup
	wg.Add(len(tests))
	for _, test := range tests {
		go wgPrintIt(&wg, test.input)
	}
	wg.Wait()

	// Cleanup: Restore stdout
	_ = w.Close()
	rawResult, _ := io.ReadAll(r)
	os.Stdout = stdOut // Restore stdout

	// Check
	result := string(rawResult)
	for _, test := range tests {
		if !strings.Contains(result, test.expected) {
			t.Errorf("Expected to find %q but it was not found", test.expected)
		}
	}
}
