package main

import (
	"testing"
	"time"

	"snippetbox.dev/internal/assert"
)

// NOTE: https://pkg.go.dev/fmt#pkg-overview
// The %q verb prints numbers and strings in quotes to express actual quotation
func TestFriendlyDate(t *testing.T) {

	// CET - Central European Time (UTC + 1)
	cetTimeZone := time.FixedZone("CET", 1*60*60)

	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "UTC",
			input:    time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC),
			expected: "03 Feb 2001 at 04:05",
		},
		{
			name:     "Zero time",
			input:    time.Time{},
			expected: "",
		},
		// [CET]04:05 is [UTC]03:05
		// Input is expressed according to CET timezone
		// Output is expected as UTC timezone
		{
			name:     "CET",
			input:    time.Date(2001, 2, 3, 4, 5, 6, 7, cetTimeZone),
			expected: "03 Feb 2001 at 03:05",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := friendlyDate(testCase.input)
			assert.Equal(t, result, testCase.expected)
		})
	}
}
