package models

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func newTestDB(t *testing.T) *sql.DB {

	// TODO: Move?
	var (
		username = "snippetboxtest"
		password = "snippetboxtest"
		host     = "localhost"
		port     = "5432"
		dbname   = "snippetboxtest"
	)

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		t.Fatal(err)
	}

	// Try reaching for the database
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	// Database setup
	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	// Database teardown
	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	})

	return db
}
