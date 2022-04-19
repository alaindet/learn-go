package main

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numLists ...[]int) []int {
	result := make([]int, len(numLists))

	for _, nums := range numLists {
		result = append(result, Sum(nums))
	}

	return result
}

// A "tail" is just a slice without the first value here
func SumAllTails(numsList ...[]int) []int {
	result := make([]int, len(numsList))

	for _, nums := range numsList {
		if len(nums) == 0 {
			result = append(result, 0)
			continue
		}

		tail := nums[1:]
		result = append(result, Sum(tail))
	}

	return result
}

// No func main here!
