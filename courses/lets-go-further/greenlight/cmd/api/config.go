package main

import (
	"flag"
	"fmt"
)

// TODO: Automate this
const version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string
}

func NewConfig() *config {

	var cfg config

	// Port
	flag.IntVar(
		&cfg.port,
		"port",
		4000, // Default
		"API server port",
	)

	// Environment
	flag.StringVar(
		&cfg.env,
		"env",
		"development", // Default
		"Environment (development|staging|production)",
	)

	// Database Source Name
	flag.StringVar(
		&cfg.dsn,
		"db-dsn",
		getDevelopmentDsn(),
		"PostgreSQL DSN",
	)

	flag.Parse()

	return &cfg
}

func getDevelopmentDsn() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		"greenlight", // username
		"greenlight", // password
		"localhost",  // host
		"5432",       // port
		"greenlight", // database name
	)
}
