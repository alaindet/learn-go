package main

import (
	"common/json"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type App struct {
	json.HTTPClient
}

func main() {
	log.Printf("Starting broker service on port %s\n", webPort)

	app := App{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
