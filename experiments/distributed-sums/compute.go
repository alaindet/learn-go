package main

func compute(c chan int, onDone func(), nums []int) {
	defer onDone()
	c <- sum(nums)
}
