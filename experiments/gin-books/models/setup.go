package models

import "gin_books/app"

func Setup() {
	app.GetApp().Database.AutoMigrate(
		&Book{},
	)
}
