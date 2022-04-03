package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {

	// Input
	csvFilename := flag.String(
		"csv",          // Flag name
		"problems.csv", // Default value
		"a csv file in the format of question,answer", // Description
	)
	timeLimit := flag.Int(
		"limit",
		30,
		"the time limit for the whole quiz in seconds",
	)
	flag.Parse()

	// Setup game
	problems := loadProblems(csvFilename)
	score := 0
	answerCh := make(chan string)
	problemsCount := len(problems)
	problemsDigits := countDigits(problemsCount)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Play
	for i, problem := range problems {
		counter := padLeft(strconv.Itoa(i+1), problemsDigits, "0")
		fmt.Printf("Problem %s/%d: %s = ", counter, problemsCount, problem.question)

		// Listen to user input concurrently
		go func(ch chan<- string) {
			var answer string
			fmt.Scanf("%s", &answer)
			ch <- answer
		}(answerCh)

		select {
		case <-timer.C: // Timeout
			fmt.Printf("\nTime's up. You score is %d out of %d\n", score, problemsCount)
			return
		case answer := <-answerCh: // Process answer
			if answer != problem.answer {
				fmt.Println("Nope")
			} else {
				fmt.Println("Yep")
				score++
			}
		}
	}

	// Result
	if score == problemsCount {
		fmt.Printf("PERFECT! You guessed all %d questions!\n", problemsCount)
		return
	}

	fmt.Printf("You score is %d out of %d\n", score, problemsCount)
}
