package core

import (
	"os"
	"strconv"
)

// Reads an environment variable as is
func env(key string, defaultVal string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	return val
}

// Reads an environment variable and converts it to an integer number
func envInt(key string, defaultVal int) int {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	intVal, _ := strconv.Atoi(val)
	return intVal
}

// Reads an environment variable and converts it to an decimal number
func envFloat(key string, defaultVal float64) float64 {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	floatVal, _ := strconv.ParseFloat(val, 64)
	return floatVal
}
