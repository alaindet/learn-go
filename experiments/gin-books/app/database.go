package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database *gorm.DB

func ConnectDatabase(s *Store) (Database, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		s.Get("DATABASE_USER"),
		s.Get("DATABASE_PASSWORD"),
		s.Get("DATABASE_HOST"),
		s.Get("DATABASE_PORT"),
		s.Get("DATABASE_NAME"),
	))

	if err != nil {
		return nil, err
	}

	return db, nil
}
