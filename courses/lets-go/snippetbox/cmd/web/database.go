package main

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
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

// TODO: This is only used for alexedwards/scs
// TODO: Merge this into openDB
func openPgxDB(dsn string) (*pgxpool.Pool, error) {
	pgxDb, err := pgxpool.Connect(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	return pgxDb, nil
}
