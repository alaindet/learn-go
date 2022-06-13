package data

import (
	"database/sql"
)

func TestNew(dbPool *sql.DB) Models {
	// Override global variable
	db = dbPool

	return Models{
		User: &TestUser{},
		Plan: &TestPlan{},
	}
}
