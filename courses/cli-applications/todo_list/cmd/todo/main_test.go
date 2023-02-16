package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = "todos.json"
)

func TestMain(m *testing.M) {

	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	buildCmd := exec.Command("go", "build", "-o", binName, ".")
	// err := buildCmd.Run()
	out, err := buildCmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s\n%s", binName, err, out)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {

	newTodo1 := "test todo number 1"
	newTodo2 := "test todo number 2"

	dir, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTodo", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", newTodo1)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AddNewTodoFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()

		if err != nil {
			t.Fatal(err)
		}

		io.WriteString(cmdStdIn, newTodo2)
		cmdStdIn.Close()
		err = cmd.Run()

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTodos", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		stdout, stderr := cmd.CombinedOutput()

		if stderr != nil {
			t.Fatal(err)
		}

		result := string(stdout)
		expected := fmt.Sprintf(
			multiline([]string{
				"[ ] %s",
				"[ ] %s",
			}),
			newTodo1,
			newTodo2,
		)

		if expected != result {
			t.Errorf("Expected %q, got %q instead\n", expected, result)
		}
	})

	t.Run("CompleteTodo", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "0")
		stdout, stderr := cmd.CombinedOutput()

		if stderr != nil {
			t.Fatal(err)
		}

		result := string(stdout)
		expected := fmt.Sprintf(
			multiline([]string{
				"Todo #0 completed.",
				"",
				"[x] %s",
				"[ ] %s",
			}),
			newTodo1,
			newTodo2,
		)

		if expected != result {
			t.Errorf("Expected %q, got %q instead\n", expected, result)
		}
	})
}

func multiline(lines []string) string {
	return strings.Join(lines, "\n") + "\n"
}
