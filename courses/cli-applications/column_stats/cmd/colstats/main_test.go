package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name        string
		col         int
		operation   string
		expected    string
		files       []string
		expectedErr error
	}{
		{
			name:        "RunAvg1File",
			col:         2,
			operation:   "avg",
			expected:    "227.6\n",
			files:       []string{"./testdata/example.csv"},
			expectedErr: nil,
		},
		{
			name:        "RunAvgMultiFiles",
			col:         2,
			operation:   "avg",
			expected:    "233.84\n",
			files:       []string{"./testdata/example.csv", "./testdata/example2.csv"},
			expectedErr: nil,
		},
		{
			name:        "RunFailRead",
			col:         1,
			operation:   "avg",
			expected:    "",
			files:       []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
			expectedErr: os.ErrNotExist,
		},
		{
			name:        "RunFailColumn",
			col:         -1,
			operation:   "avg",
			expected:    "",
			files:       []string{"./testdata/example.csv"},
			expectedErr: ErrInvalidColumn,
		},
		{
			name:        "RunFailNoFiles",
			col:         1,
			operation:   "avg",
			expected:    "",
			files:       []string{},
			expectedErr: ErrNoFiles,
		},
		{
			name:        "RunFailOperation",
			col:         1,
			operation:   "invalid",
			expected:    "",
			files:       []string{"./testdata/example.csv"},
			expectedErr: ErrInvalidOperation,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			var res bytes.Buffer
			err := run(tc.files, tc.operation, tc.col, &res)

			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("Expected error. Got nil instead")
				}
				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("Expected error %q, got %q instead", tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}

			if res.String() != tc.expected {
				t.Errorf("Expected %q, got %q instead", tc.expected, &res)
			}
		})
	}
}
