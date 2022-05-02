package main

import (
	"os"
	"time"

	"learn_go_with_tests/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
