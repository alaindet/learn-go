package main

func createRange(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, i+1)
	}
	return result
}
