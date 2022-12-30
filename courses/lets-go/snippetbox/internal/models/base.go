package models

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
)

type BaseModel struct {
	DB *sql.DB
}
