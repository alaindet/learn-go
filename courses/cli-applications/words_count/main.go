package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type counterMode string

var bytesMode counterMode = "count-bytes"
var linesMode counterMode = "count-lines"
var wordsMode counterMode = "count-words"

func main() {
	countBytes := flag.Bool("b", false, "Count bytes")
	countLines := flag.Bool("l", false, "Count lines")
	flag.Parse()

	switch {
	case *countBytes:
		fmt.Println(count(os.Stdin, bytesMode))
	case *countLines:
		fmt.Println(count(os.Stdin, linesMode))
	default:
		fmt.Println(count(os.Stdin, wordsMode))
	}
}

func count(r io.Reader, mode counterMode) int {

	scanner := bufio.NewScanner(r)

	switch mode {
	case wordsMode:
		scanner.Split(bufio.ScanWords)
	case bytesMode:
		scanner.Split(bufio.ScanBytes)
	}

	counter := 0

	for scanner.Scan() {
		counter++
	}

	return counter
}
