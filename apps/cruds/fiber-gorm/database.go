package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDsn(c config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		c.DATABASE_USER,
		c.DATABASE_PASSWORD,
		"127.0.0.1:3306", // TODO?
		c.DATABASE_NAME,
	)
}

func connectToDatabase(c config) (*sql.DB, error) {
	db, err := sql.Open("mysql", getDsn(c))
	return db, err
}
