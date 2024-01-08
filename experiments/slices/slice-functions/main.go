package main

func sliceMap[T any, K any](slice []T, mapper func(item T, i int) K) []K {
	result := make([]K, 0, len(slice))

	for i, item := range slice {
		result = append(result, mapper(item, i))
	}

	return result
}

func sliceFindIndex[T any](slice []T, finder func(item T, i int) bool) int {
	for i, item := range slice {
		if finder(item, i) {
			return i
		}
	}

	return -1
}

func sliceFind[T any](slice []T, finder func(item T, i int) bool) (T, bool) {
	for i, item := range slice {
		if finder(item, i) {
			return item, true
		}
	}

	var tZeroValue T
	return tZeroValue, false
}

func sliceFilter[T any](slice []T, filterer func(item T, i int) bool) []T {
	result := make([]T, 0, len(slice))

	for i, item := range slice {
		if filterer(item, i) {
			result = append(result, item)
		}
	}

	return result
}
