package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// TODO: Move to CLI flags/environment variables
const webPort = "80"

var dbAttempts int

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Printf("Starting authentication service on port %s\n", webPort)

	db := connectToDB()
	if db == nil {
		log.Panic("Cannot connect to database")
	}

	app := Config{
		DB:     db,
		Models: data.New(db),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
