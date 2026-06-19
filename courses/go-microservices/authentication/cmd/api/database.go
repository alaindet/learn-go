package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var dbAttempts int

// TODO: Move to CLI flags/environment variables
const dbMaxAttempts = 10
const dbBackoffSeconds = 2

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Tries to connect to server every N seconds for MAX times, then drops
// This is useful if this app starts before its database instance
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		// Success
		if err == nil {
			log.Println("Connected to database")
			return connection
		}

		// Failure
		log.Println("Database not yet ready...")
		dbAttempts++

		// Max failures reached?
		if dbAttempts > dbMaxAttempts {
			log.Println(err)
			return nil
		}

		// Trying again in a few seconds
		log.Printf("Backing off for %d seconds\n", dbBackoffSeconds)
		time.Sleep(dbBackoffSeconds * time.Second)
		continue
	}
}
