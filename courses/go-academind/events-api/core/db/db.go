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
