package main

func Map[T any, K any](elements []T, fn func(element T) K) []K {
	result := make([]K, 0, len(elements))

	for _, element := range elements {
		result = append(result, fn(element))
	}

	return result
}
