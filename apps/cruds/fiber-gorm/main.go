package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type todoApp struct {
	config config
	db     *sql.DB
}

var appData ciccio = ciccio{}

func main() {

	// Init config
	config := loadConfig()
	appData.config = config

	// Init db
	db, err := connectToDatabase(config)

	if err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	appData.db = db

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,  // "/Foo" is NOT equal to "/foo"
		StrictRouting: false, // "/foo/" is equal to "/foo"
		ServerHeader:  "Fiber",
		AppName:       c.APP_NAME,
		ReadTimeout:   time.Second * 10,
		WriteTimeout:  time.Second * 10,
	})

	app = setupRoutes(app, c)
	app.Listen(fmt.Sprintf(":%s", c.PORT))
}
