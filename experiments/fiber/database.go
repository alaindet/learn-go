package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type databaseConfig struct {
	username     string
	password     string
	host         string
	port         string
	databaseName string
}

func connectToDatabase(c *databaseConfig) (*sql.DB, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		c.username,
		c.password,
		c.host+":"+c.port,
		c.databaseName,
	)

	db, err := sql.Open("mysql", dsn)

	return db, err
}
