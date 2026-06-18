package main

import (
	"authentication/data"
	"fmt"
	"log"
	"net/http"
)

// TODO: Move to CLI flags/environment variables
const webPort = "80"

func main() {
	log.Printf("Starting authentication service on port %s\n", webPort)

	db := connectToDB()
	if db == nil {
		log.Panic("Cannot connect to database")
	}

	app := App{
		DB:     db,
		Models: data.New(db),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
