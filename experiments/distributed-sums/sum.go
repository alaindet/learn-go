package main

func sum(nums []int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}
