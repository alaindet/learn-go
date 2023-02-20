package main

// https://go.dev/src/math/all_test.go
func tolerance(a, b, e float64) bool {

	if a == b {
		return true
	}

	d := a - b
	if d < 0 {
		d = -d
	}

	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}

	return d < e
}

// https://go.dev/src/math/all_test.go
func close(a, b float64) bool { return tolerance(a, b, 1e-14) }

// https://go.dev/src/math/all_test.go
// func veryclose(a, b float64) bool { return tolerance(a, b, 4e-16) }
