package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func poolExample() {

	cfg := getDBConfig()
	connURL := getDBConnectionURL(cfg)

	// dbPool, err := pgxpool.New(context.Background(), connURL)
	dbPool, err := pgxpool.Connect(context.Background(), connURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbPool.Close()

	query := "SELECT title FROM snippets"
	var result interface{}

	err = dbPool.QueryRow(context.Background(), query).Scan(&result)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
