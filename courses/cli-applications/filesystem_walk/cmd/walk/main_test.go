package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{
			name:     "NoFilter",
			root:     "testdata",
			cfg:      config{ext: "", size: 0, list: true},
			expected: "testdata/dir.log\ntestdata/dir2/script.sh\n",
		},
		{
			name:     "FilterExtensionMatch",
			root:     "testdata",
			cfg:      config{ext: ".log", size: 0, list: true},
			expected: "testdata/dir.log\n",
		},
		{
			name:     "FilterExtensionSizeMatch",
			root:     "testdata",
			cfg:      config{ext: ".log", size: 10, list: true},
			expected: "testdata/dir.log\n",
		},
		{
			name:     "FilterExtensionSizeNoMatch",
			root:     "testdata",
			cfg:      config{ext: ".log", size: 20, list: true},
			expected: "",
		},
		{
			name:     "FilterExtensionNoMatch",
			root:     "testdata",
			cfg:      config{ext: ".gz", size: 0, list: true},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}

func TestRunDelExtension(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config
		extNoDelete string
		nDelete     int
		nNoDelete   int
		expected    string
	}{
		{
			name:        "DeleteExtensionNoMatch",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: ".gz",
			nDelete:     0,
			nNoDelete:   10,
			expected:    "",
		},
		{
			name:        "DeleteExtensionMatch",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: "",
			nDelete:     10,
			nNoDelete:   0,
			expected:    "",
		},
		{
			name:        "DeleteExtensionMixed",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: ".gz",
			nDelete:     5,
			nNoDelete:   5,
			expected:    "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var (
				buffer    bytes.Buffer
				logBuffer bytes.Buffer
			)

			tc.cfg.logFile = &logBuffer

			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext:     tc.nDelete,
				tc.extNoDelete: tc.nNoDelete,
			})
			defer cleanup()

			err := run(tempDir, &buffer, tc.cfg)
			if err != nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}

			filesLeft, err := os.ReadDir(tempDir)
			if err != nil {
				t.Error(err)
			}

			if len(filesLeft) != tc.nNoDelete {
				t.Errorf(
					"Expected %d files left, got %d instead\n",
					tc.nNoDelete,
					len(filesLeft),
				)
			}

			expectedLogLines := tc.nDelete + 1
			logLines := bytes.Split(logBuffer.Bytes(), []byte("\n"))

			if len(logLines) != expectedLogLines {
				t.Errorf(
					"Expected %d log lines, got %d instead\n",
					expectedLogLines,
					len(logLines),
				)
			}
		})
	}
}

//	createTempDir(t, map[string]int{".log": 2, ".gz": 3})
//
// Outputs:
// /TEMPDIR/file0.log
// /TEMPDIR/file1.log
// /TEMPDIR/file0.gz
// /TEMPDIR/file1.gz
// /TEMPDIR/file2.gz
func createTempDir(t *testing.T, files map[string]int) (string, func()) {

	t.Helper()
	tempDir, err := os.MkdirTemp("", "walktest")
	if err != nil {
		t.Fatal(err)
	}

	for ext, count := range files {
		for i := 1; i <= count; i++ {
			fname := fmt.Sprintf("file%d%s", i, ext)
			fpath := filepath.Join(tempDir, fname)
			err := os.WriteFile(fpath, []byte("dummy"), 0644)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	return tempDir, cleanup
}

func TestRunArchive(t *testing.T) {
	testCases := []struct {
		name         string
		cfg          config
		extNoArchive string
		nArchive     int
		nNoArchive   int
	}{
		{
			name:         "ArchiveExtensionNoMatch",
			cfg:          config{ext: ".log"},
			extNoArchive: ".gz",
			nArchive:     0,
			nNoArchive:   10,
		},
		{
			name:         "ArchiveExtensionMatch",
			cfg:          config{ext: ".log"},
			extNoArchive: "",
			nArchive:     10,
			nNoArchive:   0,
		},
		{
			name:         "ArchiveExtensionMixed",
			cfg:          config{ext: ".log"},
			extNoArchive: ".gz",
			nArchive:     5,
			nNoArchive:   5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext:      tc.nArchive,
				tc.extNoArchive: tc.nNoArchive,
			})
			defer cleanup()

			archiveDir, cleanupArchive := createTempDir(t, nil)
			defer cleanupArchive()

			tc.cfg.archive = archiveDir
			err := run(tempDir, &buffer, tc.cfg)
			if err != nil {
				t.Fatal(err)
			}

			pattern := filepath.Join(tempDir, fmt.Sprintf("*%s", tc.cfg.ext))
			expectedFiles, err := filepath.Glob(pattern)
			if err != nil {
				t.Fatal(err)
			}

			expectedOutput := strings.Join(expectedFiles, "\n")
			res := strings.TrimSpace(buffer.String())
			if expectedOutput != res {
				t.Errorf("Expected %q, got %q instead\n", expectedOutput, res)
			}

			filesArchived, err := ioutil.ReadDir(archiveDir)
			if err != nil {
				t.Fatal(err)
			}

			if len(filesArchived) != tc.nArchive {
				t.Errorf(
					"Expected %d files archived, got %d instead\n",
					tc.nArchive,
					len(filesArchived),
				)
			}
		})
	}
}
