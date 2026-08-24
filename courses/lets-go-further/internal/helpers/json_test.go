package helpers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadJSON(t *testing.T) {
	type dataType struct {
		Title string `json:"title"`
	}

	var testMaxSize int64 = 64

	testCases := []struct {
		name        string
		input       string
		shouldError bool
		maxBytes    *int64
	}{
		{
			name:        "xml",
			input:       `<?xml version="1.0" encoding="UTF-8"?><foo>Bar</foo>`,
			shouldError: true,
			maxBytes:    nil,
		},
		{
			name:        "traling comma",
			input:       `{"title": "The Title",}`,
			shouldError: true,
			maxBytes:    nil,
		},
		{
			name:        "array",
			input:       `["foo", "bar"]`,
			shouldError: true,
			maxBytes:    nil,
		},
		{
			name:        "invalid type",
			input:       `{"title": 123}`,
			shouldError: true,
			maxBytes:    nil,
		},
		{
			name:        "empty",
			input:       ``,
			shouldError: true,
			maxBytes:    nil,
		},
		{
			name:        "invalid max size",
			input:       `{"title":"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"}`,
			shouldError: true,
			maxBytes:    &testMaxSize,
		},
		{
			name:        "valid",
			input:       `{"title":"The Title"}`,
			shouldError: false,
			maxBytes:    nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			content := strings.NewReader(testCase.input)
			req := httptest.NewRequest(http.MethodPost, "/", content)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			var dest dataType
			err := ReadJSON(rec, req, &dest, testCase.maxBytes)

			if testCase.shouldError && err == nil {
				t.Errorf("Expected error, none given")
			}

			if !testCase.shouldError && err != nil {
				t.Errorf("Unexpected error")
			}
		})
	}
}
