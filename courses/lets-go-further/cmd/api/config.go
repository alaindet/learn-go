package main

import (
	"flag"
	"time"

	"app/internal/helpers/readenv"
)

const version = "1.0.0"

type config struct {
	port     int
	env      string
	database databaseConfig
}

type databaseConfig struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func readConfigFromFlags() config {
	var cfg config

	// Port
	flag.IntVar(
		&cfg.port,
		"port",
		readenv.Int("GREENLIGHT_PORT", 4000),
		"API server port",
	)

	// Environment
	flag.StringVar(
		&cfg.env,
		"env",
		readenv.String("GREENLIGHT_ENV", "development"),
		"Environment (development|staging|production)",
	)

	// Database DSN
	flag.StringVar(
		&cfg.database.dsn,
		"db-dsn",
		readenv.String("GREENLIGHT_DB_DSN", ""),
		"PostgreSQL DSN",
	)

	// Database max open connections
	flag.IntVar(
		&cfg.database.maxOpenConns,
		"db-max-open-conns",
		readenv.Int("GREENLIGHT_DB_MAX_OPEN_CONNS", 25),
		"PostgreSQL max open connections",
	)

	// Database max idle connections
	flag.IntVar(
		&cfg.database.maxIdleConns,
		"db-max-idle-conns",
		readenv.Int("GREENLIGHT_DB_MAX_IDLE_CONNS", 25),
		"PostgreSQL max idle connections",
	)

	// Database max idle time
	flag.DurationVar(
		&cfg.database.maxIdleTime,
		"db-max-idle-time",
		readenv.Duration("GREENLIGHT_DB_MAX_IDLE_TIME", 15*time.Minute),
		"PostgreSQL max connection idle time",
	)

	flag.Parse()
	return cfg
}
