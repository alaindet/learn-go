package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (app *application) openDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", app.config.dsn)

	if err != nil {
		return nil, err
	}

	// Try reaching for the database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
