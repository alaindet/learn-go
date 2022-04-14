package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type AppContext struct {
	config *config
	db     *sql.DB
	fiber  *fiber.App
}

var app *AppContext

func (a *AppContext) initConfig(envPath string) error {
	viper.SetConfigFile(envPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		return err
	}

	var c config
	viper.Unmarshal(&c)

	a.config = &c

	return nil
}

func (a *AppContext) initDatabaseConnection() error {

	dbConfig := &databaseConfig{
		username:     a.config.DATABASE_USER,
		password:     a.config.DATABASE_PASSWORD,
		host:         a.config.DATABASE_HOST,
		port:         a.config.DATABASE_PORT,
		databaseName: a.config.DATABASE_NAME,
	}

	db, err := connectToDatabase(dbConfig)

	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *AppContext) initFiberApp() error {

	// Create Fiber app
	a.fiber = fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,  // "/Foo" is NOT equal to "/foo"
		StrictRouting: false, // "/foo/" is equal to "/foo"
		ServerHeader:  "Fiber",
		AppName:       a.config.APP_NAME,
		ReadTimeout:   time.Second * 10,
		WriteTimeout:  time.Second * 10,
	})

	// Shutdown
	a.fiber.Hooks().OnShutdown(func() error {
		a.db.Close()
		log.Println("Shutting down web server...")
		return nil
	})

	// Serve static files
	a.fiber.Static("/assets", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     false, // No file streaming
		Browse:        false, // No directory browsing
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	// Register all routes
	routePrefix := fmt.Sprintf("/api/v%s", a.config.VERSION)
	a.routes(routePrefix)

	return nil
}

func initApp() error {

	app = &AppContext{}

	err := app.initConfig(".env")
	if err != nil {
		log.Fatal("Cannot load .env file")
		return err
	}

	err = app.initDatabaseConnection()
	if err != nil {
		log.Fatal("Cannot connect to database")
		return err
	}

	err = app.initFiberApp()
	if err != nil {
		log.Fatal("Cannot create Fiber app")
		return err
	}

	return nil
}
