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

	testCases := []struct {
		name        string
		input       string
		shouldError bool
	}{
		{
			name:        "xml",
			input:       `<?xml version="1.0" encoding="UTF-8"?><foo>Bar</foo>`,
			shouldError: true,
		},
		{
			name:        "traling comma",
			input:       `{"title": "The Title",}`,
			shouldError: true,
		},
		{
			name:        "array",
			input:       `["foo", "bar"]`,
			shouldError: true,
		},
		{
			name:        "invalid type",
			input:       `{"title": 123}`,
			shouldError: true,
		},
		{
			name:        "empty",
			input:       ``,
			shouldError: true,
		},
		{
			name:        "valid",
			input:       `{"title":"The Title"}`,
			shouldError: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			content := strings.NewReader(testCase.input)
			req := httptest.NewRequest(http.MethodPost, "/", content)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			_, err := ReadJSON[dataType](rec, req)

			if testCase.shouldError && err == nil {
				t.Errorf("Expected error, none given")
			}

			if !testCase.shouldError && err != nil {
				t.Errorf("Unexpected error")
			}
		})
	}
}
