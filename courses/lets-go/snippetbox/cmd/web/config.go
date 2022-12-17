package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type config struct {

	// General
	name string
	addr string

	// Paths
	staticPath        string
	htmlTemplatesPath string

	// Database
	dbUsername string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	dsn        string
}

func NewConfig() *config {
	var cfg config
	cfg.loadEnvFile()
	cfg.loadDefaultsFromEnv()
	cfg.overrideWithFlags()
	return &cfg
}

func (c *config) loadEnvFile() {
	viper.SetConfigFile(".env")
	viper.AllowEmptyEnv(true)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func (c *config) loadDefaultsFromEnv() {
	c.name = viper.GetString("SNIPPETBOX_NAME")
	c.addr = viper.GetString("SNIPPETBOX_ADDRESS")
	c.staticPath = viper.GetString("SNIPPETBOX_STATIC_PATH")
	c.htmlTemplatesPath = viper.GetString("SNIPPETBOX_HTML_TEMPLATES_PATH")
	c.dbUsername = viper.GetString("SNIPPETBOX_DB_USERNAME")
	c.dbPassword = viper.GetString("SNIPPETBOX_DB_PASSWORD")
	c.dbHost = viper.GetString("SNIPPETBOX_DB_HOST")
	c.dbPort = viper.GetString("SNIPPETBOX_DB_PORT")
	c.dbName = viper.GetString("SNIPPETBOX_DB_NAME")
	c.dsn = c.getDsn()
}

func (c *config) getDsn() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.dbUsername,
		c.dbPassword,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}

func (c *config) overrideWithFlags() {

	// Override address?
	flag.StringVar(&c.addr, "addr", c.addr,
		"HTTP network address",
	)

	// Override static path?
	flag.StringVar(&c.staticPath, "static-path", c.staticPath,
		"Path to static assets",
	)

	// Override database DSN?
	flag.StringVar(&c.dsn, "dsn", c.dsn,
		"PostgreSQL database source name",
	)

	flag.Parse()
}
