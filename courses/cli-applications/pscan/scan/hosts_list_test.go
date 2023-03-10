package scan_test

import (
	"errors"
	"os"
	"pscan/scan"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name        string
		inputHost   string
		expectedLen int
		expectedErr error
	}{
		{
			name:        "AddNew",
			inputHost:   "host2",
			expectedLen: 2,
			expectedErr: nil},
		{
			name:        "AddExisting",
			inputHost:   "host1",
			expectedLen: 1,
			expectedErr: scan.ErrExists,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			// Initialize list
			err := hl.Add("host1")
			if err != nil {
				t.Fatal(err)
			}

			err = hl.Add(tc.inputHost)

			if tc.expectedErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got 'nil' instead\n")
				}
				if !errors.Is(err, tc.expectedErr) {
					msg := "Expected error %q, got %q instead\n"
					t.Errorf(msg, tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectedLen {
				msg := "Expected list length %d, got %d instead\n"
				t.Errorf(msg, tc.expectedLen, len(hl.Hosts))
			}

			if hl.Hosts[1] != tc.inputHost {
				msg := "Expected host name %q as index 1, got %q instead\n"
				t.Errorf(msg, tc.inputHost, hl.Hosts[1])
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name        string
		inputHost   string
		expectedLen int
		expectedErr error
	}{
		{
			name:        "RemoveExisting",
			inputHost:   "host1",
			expectedLen: 1,
			expectedErr: nil,
		},
		{
			name:        "RemoveNotFound",
			inputHost:   "host3",
			expectedLen: 1,
			expectedErr: scan.ErrNotExists,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			hl := &scan.HostsList{}

			// Initialize list
			for _, h := range []string{"host1", "host2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}

			err := hl.Remove(tc.inputHost)
			if tc.expectedErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expectedErr) {
					msg := "Expected error %q, got %q instead\n"
					t.Errorf(msg, tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectedLen {
				msg := "Expected list length %d, got %d instead\n"
				t.Errorf(msg, tc.expectedLen, len(hl.Hosts))
			}

			if hl.Hosts[0] == tc.inputHost {
				t.Errorf("Host name %q should not be in the list\n", tc.inputHost)
			}
		})
	}
}

func TestSaveLoad(t *testing.T) {
	hosts1 := scan.HostsList{}
	hosts2 := scan.HostsList{}
	hosts1.Add("host1")

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	err = hosts1.Save(tf.Name())
	if err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	err = hosts2.Load(tf.Name())
	if err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	host1 := hosts1.Hosts[0]
	host2 := hosts2.Hosts[0]
	if host1 != host2 {
		t.Errorf("Host %q should match %q host.", host1, host2)
	}
}

func TestLoadNoFile(t *testing.T) {
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	err = os.Remove(tf.Name())
	if err != nil {
		t.Fatalf("Error deleting temp file: %s", err)
	}

	hl := &scan.HostsList{}
	err = hl.Load(tf.Name())
	if err != nil {
		t.Errorf("Expected no error, got %q instead\n", err)
	}
}
