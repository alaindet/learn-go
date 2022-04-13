package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c := loadConfig()

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
