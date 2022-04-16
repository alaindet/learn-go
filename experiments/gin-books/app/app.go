package app

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var app *App

type App struct {
	*Store
	Database *gorm.DB
}

func init() {
	fmt.Println("Initializing app...")
	store := NewStore(initConfig())
	db, err := ConnectDatabase(store)

	if err != nil {
		log.Fatal("Cannot initialize app")
		os.Exit(1)
	}

	app = &App{
		Store:    store,
		Database: db,
	}
}

func GetApp() *App {
	return app
}

func Get(name string) interface{} {
	return app.Store.Get(name)
}

func MustGet(name string) (interface{}, error) {
	return app.Store.MustGet(name)
}

func Set(name string, item interface{}) {
	app.Store.Set(name, item)
}
