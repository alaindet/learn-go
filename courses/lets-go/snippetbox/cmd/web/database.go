package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
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

func openPgxDB(dsn string) (*pgxpool.Pool, error) {
	pgxDb, err := pgx.Connect(dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

}
