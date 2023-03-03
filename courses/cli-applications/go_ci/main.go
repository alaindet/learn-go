package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type config struct {
	projectDir string
	out        io.Writer
}

func main() {
	cfg := parseInput()
	cfg.out = os.Stdout
	err := run(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cfg config) error {
	if cfg.projectDir == "" {
		return fmt.Errorf("project directory is required")
	}

	// Building multiple packages does not produce an output file (???)
	// Here, the given project (so the "main" package) and the "errors" package
	// Are built together
	args := []string{"build", ".", "errors"}
	cmd := exec.Command("go", args...)
	cmd.Dir = cfg.projectDir

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("'go build' failed: %s", err)
	}

	_, err = fmt.Fprintln(cfg.out, "Go Build: SUCCESS")
	return err
}

func parseInput() config {
	projectDir := flag.String("p", "", "Go project directory")
	flag.Parse()

	return config{
		projectDir: *projectDir,
		out:        nil,
	}
}
