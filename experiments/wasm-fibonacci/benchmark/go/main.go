package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	n, _ := strconv.Atoi(os.Getenv("N"))
	iterations, _ := strconv.ParseInt(os.Getenv("ITERATIONS"), 10, 64)

	start := time.Now().UnixMilli()
	var i int64
	for i < iterations {
		_ = fibonacci(n)
		i++
	}
	took := time.Now().UnixMilli() - start

	fmt.Printf("Measure time: \"Go Fibonacci\" took %d ms\n", took)
}

func fibonacci(n int) uint64 {

	if n == 0 || n == 1 {
		return uint64(n)
	}

	secondLast := uint64(0)
	last := uint64(1)
	var result uint64

	for i := 2; i <= n; i++ {
		result = secondLast + last
		secondLast = last
		last = result
	}

	return result
}
