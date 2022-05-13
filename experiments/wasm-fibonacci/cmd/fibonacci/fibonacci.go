package fibonacci

func Fibonacci(n int) uint64 {

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
