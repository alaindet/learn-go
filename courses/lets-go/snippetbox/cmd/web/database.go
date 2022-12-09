package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	// Try reaching for the database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
