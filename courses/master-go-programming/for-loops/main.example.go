/////////////////////////////////
// For Loops
// Go Playground: https://play.golang.org/p/F42s3e9KUF5
/////////////////////////////////

package main

import "fmt"

func example() {

	// printing numbers from 0 to 9
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// has the same effect as a while loop in other languages
	// there is no while loop in Go
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// handling of multiple variables in a for loop
	for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
		fmt.Printf("i = %v, j = %v\n", i, j)
	}

	// infinite loop
	// sum := 0
	// for {
	//  sum++
	// }
	// fmt.Println(sum) //this line is never reached
}
