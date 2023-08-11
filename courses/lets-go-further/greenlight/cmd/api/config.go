package main

import (
	"flag"
	"os"
	"strconv"
)

// TODO: Automate this
const version = "1.0.0"

type databaseConfig struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type rateLimitConfig struct {
	avg float64
	max int
}

type config struct {
	port      int
	env       string
	db        databaseConfig
	rateLimit rateLimitConfig
}

func NewConfig() *config {

	var cfg config

	// Port
	flag.IntVar(
		&cfg.port,
		"port",
		envInt("GREENLIGHT_PORT", 4000),
		"API server port",
	)

	// Environment
	flag.StringVar(
		&cfg.env,
		"env",
		env("GREENLIGHT_ENV", "development"),
		"Environment (development|staging|production)",
	)

	// Database Source Name
	flag.StringVar(
		&cfg.db.dsn,
		"db-dsn",
		env("GREENLIGHT_DB_DSN", "postgres://greenlight:greenlight@localhost:5432/greenlight"),
		"PostgreSQL DSN",
	)

	// Database max open connections
	flag.IntVar(
		&cfg.db.maxOpenConns,
		"db-max-open-conns",
		envInt("GREENLIGHT_DB_MAX_OPEN_CONNS", 25),
		"PostgreSQL max open connections",
	)

	// Database max idle connections
	flag.IntVar(
		&cfg.db.maxIdleConns,
		"db-max-idle-conns",
		envInt("GREENLIGHT_DB_MAX_IDLE_CONNS", 25),
		"PostgreSQL max idle connections",
	)

	// Database max idle time
	flag.StringVar(
		&cfg.db.maxIdleTime,
		"db-max-idle-time",
		env("GREENLIGHT_DB_MAX_IDLE_TIME", "15m"),
		"PostgreSQL max connection idle time",
	)

	// Rate limit average
	flag.Float64Var(
		&cfg.rateLimit.avg,
		"rate-limit-global-avg",
		envFloat("GREENLIGHT_RATE_LIMIT_GLOBAL_AVG", 2.0),
		"Global rate limit average requests/second",
	)

	flag.Parse()

	return &cfg
}

func env(key string, defaultVal string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	return val
}

func envInt(key string, defaultVal int) int {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	intVal, _ := strconv.Atoi(val)
	return intVal
}

func envFloat(key string, defaultVal float64) float64 {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	floatVal, _ := strconv.ParseFloat(val, 64)
	return floatVal
}
