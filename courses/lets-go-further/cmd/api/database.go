package main

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

var (
	pingTimeout = 5 * time.Second
)

func openDB(logger *slog.Logger, cfg databaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.maxOpenConns)
	db.SetMaxIdleConns(cfg.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.maxIdleTime)

	logger.Info(
		"database set up",
		"max open connections", cfg.maxOpenConns,
		"max idle connections", cfg.maxIdleConns,
		"max idle time", cfg.maxIdleTime,
	)

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
