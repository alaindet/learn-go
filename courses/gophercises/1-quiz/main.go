package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of question,answer")
	flag.Parse()
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to parse %s\n", *csvFilename))
	}

	problems := parseLines(lines)
	score := 0
	problemsCount := len(problems)
	problemsDigits := countDigits(problemsCount)

	for i, problem := range problems {
		counter := padLeft(strconv.Itoa(i+1), problemsDigits, "0")
		fmt.Printf("Problem %s/%d: %s = ", counter, problemsCount, problem.question)
		var answer string
		fmt.Scanf("%s", &answer)

		if answer != problem.answer {
			fmt.Println("Nope")
			continue
		}

		fmt.Println("Yep")
		score++
	}

	fmt.Printf("You score is %d out of %d\n", score, len(problems))
}
