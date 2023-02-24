package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Example file
// Col1,Col2
// Data0,60707
// Data1,25641
// Data2,79731
// Data3,18485
func createCsvFile(path string, lines int) error {
	f, err := os.Create(path)

	// Cannot create file
	if err != nil {
		return err
	}

	defer f.Close()
	w := bufio.NewWriter(f)
	random := getRandomGenerator(10000, 100000)

	// Headers
	_, err = w.WriteString("Col1,Col2\n")

	if err != nil {
		return err
	}

	// Generate rows
	for i := 0; i < lines; i++ {
		line := fmt.Sprintf("Data%d,%d\n", i, random())
		_, err = w.WriteString(line)
		if err != nil {
			return err
		}
	}

	w.Flush()
	return nil
}

func getRandomGenerator(min, max int) func() int {
	rand.Seed(time.Now().UnixNano())
	return func() int {
		return rand.Intn(max-min+1) + min
	}
}
