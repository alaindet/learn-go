package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func loadProblems(csvFilename *string) []problem {
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to parse %s\n", *csvFilename))
	}

	return parseLines(lines)
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
