package models

import "database/sql"

type baseModel struct {
	DB *sql.DB
}
