package main

import (
	"common/json"
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = "80"
)

type App struct {
	json.HTTPClient
	Mailer Mail
}

func main() {
	log.Printf("Starting mail service on port %s\n", webPort)

	mailer, err := NewMail()
	if err != nil {
		log.Panic(err)
	}

	app := App{
		Mailer: mailer,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
