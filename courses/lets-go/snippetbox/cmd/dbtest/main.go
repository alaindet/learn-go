// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// TODO: Move into .env
	dbUser := "postgres"
	dbPassword := "postgres"
	dbHost := "localhost"
	dbName := "snippetbox"
	dbPort := "5432"

	connectionString := fmt.Sprintf(
		// Ex.: "postgres://username:password@localhost:5432/database_name"
		"postgres://%s:%s@%s:%s/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	conn, err := pgx.Connect(context.Background(), connectionString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var result interface{}
	query := "SELECT title FROM snippets"
	err = conn.QueryRow(context.Background(), query).Scan(&result)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("RESULT: %+v\n", result)
}
