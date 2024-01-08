package main

import (
	"math/rand"
	"testing"
)

func BenchmarkFasterConcMap1000(b *testing.B) {
	input, output := createBenchmarkVars(1_000)

	for n := 0; n < b.N; n++ {
		output = FasterConcMap(input, rand.Intn, -1)
	}
	_ = output
}

func BenchmarkConcMap1000(b *testing.B) {
	input, output := createBenchmarkVars(1_000)

	for n := 0; n < b.N; n++ {
		output = ConcMap(input, rand.Intn, -1)
	}
	_ = output
}

func BenchmarkMap1000(b *testing.B) {
	input, output := createBenchmarkVars(1_000)

	for n := 0; n < b.N; n++ {
		output = Map(input, rand.Intn)
	}
	_ = output
}

func BenchmarkFasterConcMap1000000(b *testing.B) {
	input, output := createBenchmarkVars(1_000_000)

	for n := 0; n < b.N; n++ {
		output = FasterConcMap(input, rand.Intn, -1)
	}
	_ = output
}

func BenchmarkConcMap1000000(b *testing.B) {
	input, output := createBenchmarkVars(1_000_000)

	for n := 0; n < b.N; n++ {
		output = ConcMap(input, rand.Intn, -1)
	}
	_ = output
}

func BenchmarkMap1000000(b *testing.B) {
	input, output := createBenchmarkVars(1_000_000)

	for n := 0; n < b.N; n++ {
		output = Map(input, rand.Intn)
	}
	_ = output
}

func createBenchmarkVars(size int) ([]int, []int) {
	input := createRange(size)
	output := make([]int, 0, len(input))
	return input, output
}
