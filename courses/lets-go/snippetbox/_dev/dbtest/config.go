package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DBConfig struct {
	username string
	password string
	host     string
	port     string
	name     string
}

// TODO: Get from .env
func getDBConfig() *DBConfig {
	return &DBConfig{
		username: "postgres",
		password: "postgres",
		host:     "localhost",
		port:     "5432",
		name:     "snippetbox",
	}
}

// Ex.: "postgres://username:password@localhost:5432/database_name"
func getDBConnectionURL(cfg *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.username,
		cfg.password,
		cfg.host,
		cfg.port,
		cfg.name,
	)
}
