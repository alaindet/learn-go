package utils

import "strings"

func IntsToBools(ints []int) []bool {
	result := make([]bool, len(ints))
	for i, val := range ints {
		if val == 0 {
			result[i] = false
		} else {
			result[i] = true
		}
	}
	return result
}

func JoinLines(lines ...string) string {
	return strings.Join(lines, "\n")
}
