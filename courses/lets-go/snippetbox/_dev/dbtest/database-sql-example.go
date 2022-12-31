package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func databaseSqlExample() {

	cfg := getDBConfig()
	connURL := getDBConnectionURL(cfg)

	db, err := sql.Open("pgx", connURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	query := "SELECT title FROM snippets"
	var result interface{}
	err = db.QueryRow(query).Scan(&result)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
