package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}

	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	})
}
