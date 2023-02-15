package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

	build := exec.Command("go", "build", "-o", binName, "./cmd/todo")
	err := build.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
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
	todo := "test todo number 1"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTodo", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-todo", todo)
		err := cmd.Run()
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
		expected := todo + "\n"

		if expected != result {
			t.Errorf("Expected %q, got %q instead\n", expected, result)
		}
	})
}
