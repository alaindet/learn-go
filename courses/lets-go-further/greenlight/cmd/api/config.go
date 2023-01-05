package main

import "flag"

// TODO: Automate this
const version = "1.0.0"

type config struct {
	port int
	env  string
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

	flag.Parse()

	return &cfg
}
