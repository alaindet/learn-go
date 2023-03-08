package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {

	// Skip all if not git command is found
	_, err := exec.LookPath("git")
	if err != nil {
		t.Skip("Git not installed. Skipping test.")
	}

	var testCases = []struct {
		name           string
		inputProject   string
		expectedOutput string
		expectedError  error
		setupGit       bool
	}{
		{
			name:           "success",
			inputProject:   "./testdata/tool/",
			expectedOutput: "Go Build: SUCCESS\nGo Test: SUCCESS\nGofmt: SUCCESS\nGit Push: SUCCESS\n",
			expectedError:  nil,
			setupGit:       true,
		},
		{
			name:           "fail",
			inputProject:   "./testdata/tool_err",
			expectedOutput: "",
			expectedError:  &stepErr{step: "go build"},
			setupGit:       false,
		},
		{
			name:           "failformat",
			inputProject:   "./testdata/tool_gofmt_err",
			expectedOutput: "",
			expectedError:  &stepErr{step: "go fmt"},
			setupGit:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.setupGit {
				cleanup := setupGit(t, tc.inputProject)
				defer cleanup()
			}

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

// Test helper
func setupGit(t *testing.T, projectDir string) func() {
	t.Helper()

	// Search for installed git command
	gitExec, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "gocitest")
	if err != nil {
		t.Fatal(err)
	}

	// Make project path absolute if not already absolute
	projectPath, err := filepath.Abs(projectDir)
	if err != nil {
		t.Fatal(err)
	}

	// Create a fake remote Git origin
	remoteURI := fmt.Sprintf("file://%s", tempDir)

	var gitCommands = []struct {
		args []string
		dir  string
		env  []string
	}{
		{[]string{"init", "--bare"}, tempDir, nil},
		{[]string{"init"}, projectDir, nil},
		{[]string{"remote", "add", "origin", remoteURI}, projectDir, nil},
		{[]string{"add", "."}, projectDir, nil},
		{[]string{"commit", "-m", "test"}, projectDir,
			[]string{
				"GIT_COMMITTER_NAME=test",
				"GIT_COMMITTER_EMAIL=test@example.com",
				"GIT_AUTHOR_NAME=test",
				"GIT_AUTHOR_EMAIL=test@example.com",
			}},
	}

	// Setup a bare Git repository for pushing only
	for _, command := range gitCommands {
		gitCmd := exec.Command(gitExec, command.args...)
		gitCmd.Dir = command.dir
		if command.env != nil {
			gitCmd.Env = append(os.Environ(), command.env...)
		}

		err := gitCmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	}

	// Cleanup
	return func() {
		os.RemoveAll(tempDir)
		os.RemoveAll(filepath.Join(projectPath, ".git"))
	}
}
