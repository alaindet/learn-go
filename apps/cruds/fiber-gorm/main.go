package main

import (
	"fmt"
	"log"
)

func main() {
	c := loadConfig()

	db, err := connectToDatabase(c)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	rows, err := db.Query("SELECT 42")

	if err != nil {
		log.Fatal("Cannot select from database")
		return
	}

	fmt.Println("rows", rows)

	// app := fiber.New(fiber.Config{
	// 	Prefork:       true,
	// 	CaseSensitive: true,  // "/Foo" is NOT equal to "/foo"
	// 	StrictRouting: false, // "/foo/" is equal to "/foo"
	// 	ServerHeader:  "Fiber",
	// 	AppName:       c.APP_NAME,
	// 	ReadTimeout:   time.Second * 10,
	// 	WriteTimeout:  time.Second * 10,
	// })

	// app = setupRoutes(app, c)
	// app.Listen(fmt.Sprintf(":%s", c.PORT))
}
