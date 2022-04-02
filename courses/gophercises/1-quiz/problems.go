package main

import "strings"

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	// This initializes a slice of problem struct with length len(lines)
	// All values are empty problem struct (zero value)
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}
