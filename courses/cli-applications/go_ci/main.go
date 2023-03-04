package main

import (
	"fmt"
	"os"
)

func main() {
	cfg := parseInput()
	cfg.out = os.Stdout

	err := run(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
