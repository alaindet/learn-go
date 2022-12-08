package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	name       string
	addr       string
	staticPath string
}

// Load from .env, override some vars with optional CLI flags
func loadConfig() *config {
	var cfg config
	loadEnvironmentFile()

	cfg.name = viper.GetString("SNIPPETBOX_NAME")

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
