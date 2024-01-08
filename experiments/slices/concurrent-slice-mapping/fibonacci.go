package main

func Fibonacci(nth int) int64 {

	if nth <= 1 {
		return int64(nth)
	}

	var a int64 = 0
	var b int64 = 1

	for i := 2; i <= nth; i++ {
		a, b = b, b+a
	}

	return b
}
