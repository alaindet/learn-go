package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
func helloWorldExample() {

	cfg := getDBConfig()
	connURL := getDBConnectionURL(cfg)

	conn, err := pgx.Connect(context.Background(), connURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	query := "SELECT title FROM snippets"
	var result interface{}

	err = conn.QueryRow(context.Background(), query).Scan(&result)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("RESULT: %+v\n", result)
}
