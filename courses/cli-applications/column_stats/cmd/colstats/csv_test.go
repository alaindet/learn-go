package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

func TestOperations(t *testing.T) {
	data := [][]float64{
		{10, 20, 15, 30, 45, 50, 100, 30},
		{5.5, 8, 2.2, 9.75, 8.45, 3, 2.5, 10.25, 4.75, 6.1, 7.67, 12.287, 5.47},
		{-10, -20},
		{102, 37, 44, 57, 67, 129},
	}

	testCases := []struct {
		name      string
		operation statsFunc
		expected  []float64
	}{
		{"Sum", sum, []float64{300, 85.927, -30, 436}},
		{"Avg", avg, []float64{37.5, 6.609769230769231, -15, 72.666666666666666}},
		{"Min", min, []float64{10, 2.2, -20, 37}},
		{"Max", max, []float64{100, 12.287, -10, 129}},
	}

	for _, tc := range testCases {
		for i, expectedVal := range tc.expected {
			name := fmt.Sprintf("%sData%d", tc.name, i)
			t.Run(name, func(t *testing.T) {
				res := tc.operation(data[i])
				if !approxEqual(res, expectedVal) {
					t.Errorf("Expected %g, got %g instead", expectedVal, res)
				}
			})
		}
	}
}

func TestCSVV2Float(t *testing.T) {
	csvData := strings.Join([]string{
		"IP Address,Requests,Response Time",
		"192.168.0.199,2056,236",
		"192.168.0.88,899,220",
		"192.168.0.199,3054,226",
		"192.168.0.100,4133,218",
		"192.168.0.199,950,238",
	}, "\n")

	testCases := []struct {
		name        string
		col         int
		expected    []float64
		expectedErr error
		r           io.Reader
	}{
		{
			name:        "Column2",
			col:         1,
			expected:    []float64{2056, 899, 3054, 4133, 950},
			expectedErr: nil,
			r:           bytes.NewBufferString(csvData),
		},
		{
			name:        "Column3",
			col:         2,
			expected:    []float64{236, 220, 226, 218, 238},
			expectedErr: nil,
			r:           bytes.NewBufferString(csvData),
		},
		{
			name:        "FailRead",
			col:         0,
			expected:    nil,
			expectedErr: iotest.ErrTimeout,
			r:           iotest.TimeoutReader(bytes.NewReader([]byte{0})),
		},
		{
			name:        "FailedNotNumber",
			col:         0,
			expected:    nil,
			expectedErr: ErrNotNumber,
			r:           bytes.NewBufferString(csvData),
		},
		{
			name:        "FailedInvalidColumn",
			col:         42,
			expected:    nil,
			expectedErr: ErrInvalidColumn,
			r:           bytes.NewBufferString(csvData),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := csv2float(tc.r, tc.col)

			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error. Got nil instead")
				}
				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("expected error %q, got %q instead", tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}

			for i, exp := range tc.expected {
				if res[i] != exp {
					t.Errorf("Expected %g, got %g instead", exp, res[i])
				}
			}
		})
	}
}
