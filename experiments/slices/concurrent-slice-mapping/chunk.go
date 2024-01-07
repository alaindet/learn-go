package main

func Chunk[T any](s []T, chunkSize int) [][]T {

	chunksCount := getChunksCount(len(s), chunkSize)
	chunks := make([][]T, 0, chunksCount)

	if chunkSize >= len(s) {
		chunks = append(chunks, s)
		return chunks
	}

	for i := 0; i < chunksCount; i++ {
		from := i * chunkSize
		to := min(from+chunkSize, len(s))
		chunks = append(chunks, s[from:to])
	}

	return chunks
}

func getChunksCount(size, chunkSize int) int {
	result := size / chunkSize

	if size%chunkSize > 0 {
		result += 1
	}

	return result
}
