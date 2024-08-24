package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "events.db")
	if err != nil {
		panic("Could not connect to the database")
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5)
	DB = db
	createTables()
}

func createTables() {
	createUsersTable()
	createEventsTable()
	createEventParticipationsTable()
}

func createUsersTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS "users" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"email" TEXT NOT NULL UNIQUE,
			"password" TEXT NOT NULL
		)
	`)

	if err != nil {
		panic("Could not create \"users\" table")
	}
}

func createEventsTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS "events" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name" TEXT NOT NULL,
			"description" TEXT NOT NULL,
			"location" TEXT NOT NULL,
			"date_time" DATETIME NOT NULL,
			"user_id" INTEGER,
			FOREIGN KEY("user_id") REFERENCES "users"("id")
		)
	`)

	if err != nil {
		panic("Could not create \"events\" table")
	}
}

func createEventParticipationsTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS "event_participations" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"event_id" INTEGER,
			"user_id" INTEGER,
			UNIQUE ("event_id", "user_id"),
			FOREIGN KEY("event_id") REFERENCES "events"("id"),
			FOREIGN KEY("user_id") REFERENCES "users"("id")
		)
	`)

	if err != nil {
		panic("Could not create \"event_participations\" table")
	}
}
