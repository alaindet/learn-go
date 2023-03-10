package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"testing"
	"time"
)

func TestRun(t *testing.T) {

	// // Skip all if not git command is found
	// _, err := exec.LookPath("git")
	// if err != nil {
	// 	t.Skip("Git not installed. Skipping test.")
	// }

	successOutput := "Go Build: SUCCESS\n" +
		"Go Test: SUCCESS\n" +
		"Gofmt: SUCCESS\n" +
		"Git Push: SUCCESS\n"

	var testCases = []struct {
		name           string
		inputProject   string
		expectedOutput string
		expectedError  error
		setupGit       bool
		mockCmd        func(ctx context.Context, name string, arg ...string) *exec.Cmd
	}{
		{
			name:           "success",
			inputProject:   "./testdata/tool/",
			expectedOutput: successOutput,
			expectedError:  nil,
			setupGit:       true,
			mockCmd:        nil,
		},
		{
			name:           "successMock",
			inputProject:   "./testdata/tool/",
			expectedOutput: successOutput,
			expectedError:  nil,
			setupGit:       false,
			mockCmd:        mockCmdContext,
		},
		{
			name:           "fail",
			inputProject:   "./testdata/tool_err",
			expectedOutput: "",
			expectedError:  &stepErr{step: "go build"},
			setupGit:       false,
			mockCmd:        nil,
		},
		{
			name:           "failformat",
			inputProject:   "./testdata/tool_gofmt_err",
			expectedOutput: "",
			expectedError:  &stepErr{step: "go fmt"},
			setupGit:       false,
			mockCmd:        nil,
		},
		{
			name:           "failTimeout",
			inputProject:   "./testdata/tool",
			expectedOutput: "",
			expectedError:  context.DeadlineExceeded,
			setupGit:       false,
			mockCmd:        mockCmdTimeout,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.setupGit {
				_, err := exec.LookPath("git")
				if err != nil {
					t.Skip("Git not installed. Skipping test.")
				}
				cleanup := setupGit(t, tc.inputProject)
				defer cleanup()
			}

			if tc.mockCmd != nil {
				command = tc.mockCmd
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

func TestRunKill(t *testing.T) {
	var testCases = []struct {
		name          string
		inputProject  string
		sig           syscall.Signal
		expectedError error
	}{
		{
			name:          "SIGINT",
			inputProject:  "./testdata/tool",
			sig:           syscall.SIGINT,
			expectedError: ErrSignal,
		},
		{
			name:          "SIGTERM",
			inputProject:  "./testdata/tool",
			sig:           syscall.SIGTERM,
			expectedError: ErrSignal,
		},
		{
			name:          "SIGQUIT",
			inputProject:  "./testdata/tool",
			sig:           syscall.SIGQUIT,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			command = mockCmdTimeout
			errCh := make(chan error)
			ignoredSigCh := make(chan os.Signal, 1)
			expectedSigCh := make(chan os.Signal, 1)

			signal.Notify(ignoredSigCh, syscall.SIGQUIT)
			defer signal.Stop(ignoredSigCh)

			signal.Notify(expectedSigCh, tc.sig)
			defer signal.Stop(expectedSigCh)

			// Run the function
			go func() {
				errCh <- run(config{
					projectDir: tc.inputProject,
					out:        io.Discard,
				})
			}()

			// Kill if after 1 second
			go func() {
				time.Sleep(1 * time.Second)
				syscall.Kill(syscall.Getpid(), tc.sig)
			}()

			// Read the signal
			select {
			case err := <-errCh:
				if err == nil {
					t.Errorf("Expected error. Got 'nil' instead.")
					return
				}
				if !errors.Is(err, tc.expectedError) {
					t.Errorf("Expected error: %q. Got %q", tc.expectedError, err)
				}
				// Read the error
				select {
				case rec := <-expectedSigCh:
					if rec != tc.sig {
						t.Errorf("Expected signal %q, got %q", tc.sig, rec)
					}
				default:
					t.Errorf("Signal not received")
				}
			case <-ignoredSigCh:
				// Do nothing
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
	// remoteURI := fmt.Sprintf("file://%s", tempDir)
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

			fmt.Println("setupGit 5", command) // TODO: Remove

			t.Fatal(err)
		}
	}

	// Cleanup
	return func() {
		os.RemoveAll(tempDir)
		os.RemoveAll(filepath.Join(projectPath, ".git"))
	}
}

func mockCmdContext(ctx context.Context, exe string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess"}
	cs = append(cs, exe)
	cs = append(cs, args...)
	// os.Args[0] is the name of the temporary binary file created to run the test
	cmd := exec.CommandContext(ctx, os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func mockCmdTimeout(ctx context.Context, exe string, args ...string) *exec.Cmd {
	cmd := mockCmdContext(ctx, exe, args...)
	cmd.Env = append(cmd.Env, "GO_HELPER_TIMEOUT=1")
	return cmd
}

func TestHelperProcess(t *testing.T) {
	// Skip execution if not called from helper
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	// Simulate a long timeout to avoid timeouts
	if os.Getenv("GO_HELPER_TIMEOUT") == "1" {
		time.Sleep(15 * time.Second)
	}

	// If the test attemps to call git, return 0 (everything's ok) but do not execute it
	if os.Args[2] == "git" {
		fmt.Fprintln(os.Stdout, "Everything up-to-date")
		os.Exit(0)
	}

	os.Exit(1)
}
