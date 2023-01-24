package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(cfg databaseConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.dsn)
	if err != nil {
		return nil, err
	}

	maxIdleTime, err := time.ParseDuration(cfg.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.maxOpenConns)
	db.SetMaxIdleConns(cfg.maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	// Try reaching the database for 5 seconds, then fail
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
