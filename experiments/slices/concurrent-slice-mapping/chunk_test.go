package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	t.Run(
		"returns one chunk is chunk size is larger than slice size",
		func(t *testing.T) {
			s := []int{1, 2, 3}
			outcome := Chunk(s, 10)
			assert.Equal(t, outcome[0], s)
		},
	)

	t.Run(
		"returns a smaller last chunk if slice length is not multiple of chunk size",
		func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			chunkSize := 2
			outcome := Chunk(s, chunkSize)
			normalChunk := outcome[0]
			lastChunk := outcome[len(outcome)-1]
			assert.Greater(t, len(normalChunk), len(lastChunk))
		},
	)

	t.Run(
		"returns equal chunks if slice length is multiple of chunk size",
		func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5, 6}
			chunkSize := 2
			outcome := Chunk(s, chunkSize)
			normalChunk := outcome[0]
			lastChunk := outcome[len(outcome)-1]
			assert.Equal(t, len(normalChunk), len(lastChunk))
		},
	)
}
