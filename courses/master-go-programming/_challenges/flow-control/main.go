package main

import (
	"fmt"
	"time"
)

func main() {
	// #1
	for i := 0; i < 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")

	// #2
	for i := 0; i < 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// #3
	count := 0
	for i := 0; i < 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)
			count++
			if count == 3 {
				break
			}
		}
	}
	fmt.Printf("\n")

	// #4
	for i := 1; i < 501; i++ {
		if i%5 == 0 || i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")

	// #5
	for year, currentYear := 1990, time.Now().Year(); year <= currentYear; year++ {
		if year != currentYear {
			fmt.Printf("%d, ", year)
		} else {
			fmt.Printf("%d\n", year)
			break
		}
	}
}
