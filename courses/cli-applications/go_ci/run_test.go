package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	var testCases = []struct {
		name           string
		inputProject   string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "success",
			inputProject:   "./testdata/tool/",
			expectedOutput: "Go Build: SUCCESS\n",
			expectedError:  nil,
		},
		{
			name:           "fail",
			inputProject:   "./testdata/tool_err",
			expectedOutput: "",
			expectedError:  &stepErr{step: "go build"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var out bytes.Buffer

			err := run(config{
				projectDir: tc.inputProject,
				out:        &out,
			})

			if tc.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error: %q. Got 'nil' instead.", tc.expectedError)
					return
				}
				if !errors.Is(err, tc.expectedError) {
					t.Errorf("Expected error: %q. Got %q.", tc.expectedError, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}

			if out.String() != tc.expectedOutput {
				t.Errorf("Expected output: %q. Got %q", tc.expectedOutput, out.String())
			}
		})
	}
}
