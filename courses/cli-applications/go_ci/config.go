package main

import (
	"flag"
	"io"
)

type config struct {
	projectDir string
	out        io.Writer
}

func parseInput() config {
	projectDir := flag.String("p", "", "Go project directory")
	flag.Parse()

	return config{
		projectDir: *projectDir,
		out:        nil,
	}
}
