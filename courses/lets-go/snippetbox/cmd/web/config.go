package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	addr       string
	staticPath string
}

// Load from .env, override with optional CLI flags
func loadConfig() *config {
	var cfg config
	loadEnvironmentFile()

	// Address
	flag.StringVar(
		&cfg.addr,
		"addr",
		viper.GetString("SNIPPETBOX_ADDRESS"),
		"HTTP network address",
	)

	// Static path
	flag.StringVar(
		&cfg.staticPath,
		"static-path",
		viper.GetString("SNIPPETBOX_STATIC_PATH"),
		"Path to static assets",
	)

	flag.Parse()
	return &cfg
}

func loadEnvironmentFile() {
	viper.SetConfigFile(".env")
	viper.AllowEmptyEnv(true)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
