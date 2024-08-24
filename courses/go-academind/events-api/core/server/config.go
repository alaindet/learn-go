package server

import "flag"

type ServerConfig struct {
	Port string
}

func ReadConfigFromCLI() ServerConfig {
	cfg := ServerConfig{}

	flag.StringVar(&cfg.Port, "port", "8080", "Server port")
	flag.Parse()

	return cfg
}
