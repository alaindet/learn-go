package main

import "fmt"

func main() {
	// #1
	var cities = [2]string{}
	_ = cities

	// #2
	grades := [3]float64{8, 9}
	fmt.Println(grades) // [6 7 0]

	// #3
	taskDone := [...]bool{true, false, true} // <- inline declarations can skip trailing comma
	fmt.Println(taskDone)

	// #4
	cities[0] = "Rome"
	cities[1] = "Milan"
	for i, I := 0, len(cities); i < I; i++ {
		city := cities[i]
		fmt.Printf("%v ", city)
	}

	// #5
	for i, grade := range grades {
		fmt.Printf("index: %d, grade: %v\n", i, grade)
	}
}
