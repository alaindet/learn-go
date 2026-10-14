package readenv

import (
	"os"
	"strconv"
	"time"
)

func Int(envKey string, defaultValue int) int {
	envValue := os.Getenv(envKey)
	if envValue == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(envValue)
	if err != nil {
		return defaultValue
	}

	return value
}

func String(envKey string, defaultValue string) string {
	envValue := os.Getenv(envKey)
	if envValue == "" {
		return defaultValue
	}

	return envValue
}

func Duration(envKey string, defaultValue time.Duration) time.Duration {
	envValue := os.Getenv(envKey)
	if envValue == "" {
		return defaultValue
	}

	// Accepted durations are numbers with specified unit suffix
	// Suffixes are: ns, us, ms, s, m, h
	value, err := time.ParseDuration(envValue)
	if err != nil {
		return defaultValue
	}

	return value
}
