package models

import "database/sql"

type baseModel struct {
	db *sql.DB
}
