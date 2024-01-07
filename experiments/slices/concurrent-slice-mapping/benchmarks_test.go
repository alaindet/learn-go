package main

import "testing"

func BenchmarkConcMap100(b *testing.B) {
	b.StopTimer()
	input := createRange(100)
	output := make([]int, 0, len(input))
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		output = ConcMap(input, double, -1)
	}
	_ = output
}

func BenchmarkMap100(b *testing.B) {
	b.StopTimer()
	input := createRange(100)
	output := make([]int, 0, len(input))
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		output = Map(input, double)
	}
	_ = output
}

func BenchmarkConcMap10000(b *testing.B) {
	b.StopTimer()
	input := createRange(10_000)
	output := make([]int, 0, len(input))
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		output = ConcMap(input, double, -1)
	}
	_ = output
}

func BenchmarkMap10000(b *testing.B) {
	b.StopTimer()
	input := createRange(10_000)
	output := make([]int, 0, len(input))
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		output = Map(input, double)
	}
	_ = output
}

func BenchmarkConcMap1000000(b *testing.B) {
	b.StopTimer()
	input := createRange(1_000_000)
	output := make([]int, 0, len(input))
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		output = ConcMap(input, double, -1)
	}
	_ = output
}

func BenchmarkMap1000000(b *testing.B) {
	b.StopTimer()
	input := createRange(1_000_000)
	output := make([]int, 0, len(input))
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		output = Map(input, double)
	}
	_ = output
}

func createRange(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, i+1)
	}
	return result
}
