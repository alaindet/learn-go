package main

import (
	"fmt"
	"log"
)

func main() {

	err := initApp()

	if err != nil {
		log.Fatal("Cannot initialize application")
		return
	}

	baseUrl := fmt.Sprintf("/api/v%s", app.config.VERSION)
	app.routes(baseUrl)

	address := fmt.Sprintf(":%s", app.config.PORT)
	app.fiber.Listen(address)
}
