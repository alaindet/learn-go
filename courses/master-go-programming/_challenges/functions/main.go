package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func cube(x float64) float64 {
	return x * x * x
}

func myFactorial(x uint) (uint, uint) {
	if x > 12 {
		return 0, 0
	}

	var sum uint = 0
	var factorial uint = 1
	var i uint = 1

	for i < x {
		factorial *= i
		sum += i
		i++
	}

	return factorial, sum
}

func parseAndSum(inputNum string) int {
	n, err := strconv.Atoi(inputNum)

	if err != nil {
		log.Fatal(err)
	}

	return 100*1*n + 10*2*n + 1*3*n
}

func mySum(operands ...int) int {
	sum := 0
	for _, value := range operands {
		sum += value
	}
	return sum
}

func mySumNaked(operands ...int) (sum int) {
	for _, value := range operands {
		sum += value
	}
	return
}

func contains(haystack []string, needle string) bool {
	for _, word := range haystack {
		if word == needle {
			return true
		}
	}

	return false
}

func containsCaseInsensitive(haystack []string, needle string) bool {
	needleLower := strings.ToLower(needle)
	for _, word := range haystack {
		if strings.ToLower(word) == needleLower {
			return true
		}
	}

	return false
}

func getInt(x int) int {
	return x
}

func main() {
	// #1
	ex1 := cube(3.0)
	fmt.Println(ex1) // 27

	// #2
	fact, sum := myFactorial(11)
	fmt.Println(fact, sum) // 3628800 55

	// #3
	ex3 := parseAndSum("5")
	fmt.Println(ex3) // 615

	// #4
	ex4 := mySum(1, 2, 3, 4)
	fmt.Println(ex4) // 10

	// #5
	ex5 := mySumNaked(1, 2, 3, 4)
	fmt.Println(ex5) // 10

	// #6
	words := []string{"Hello", "World"}
	ex6a := contains(words, "World")
	ex6b := contains(words, "Universe")
	ex6c := contains(words, "world") // Case-sensitive search!
	fmt.Println(ex6a, ex6b, ex6c)    // true false false

	// #7
	ex7a := containsCaseInsensitive(words, "World")
	ex7b := containsCaseInsensitive(words, "Universe")
	ex7c := containsCaseInsensitive(words, "world")
	fmt.Println(ex7a, ex7b, ex7c) // true false true

	// #8
	// ...

	// #9
	getter := getInt
	fmt.Printf("%T %v", getter, getter(42)) // func(int) int 42
}
