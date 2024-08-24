package models

import (
	"app/common/utils"
	"app/core/db"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrUserExists = errors.New("user already exists")
)

var createSql = `
	INSERT INTO "users" ("email", "password")
	VALUES (?, ?)
`

func (u UserModel) Create() (UserModel, error) {

	stmt, err := db.DB.Prepare(createSql)
	if err != nil {
		return u, err
	}
	defer stmt.Close()

	password, err := utils.HashPassword(u.Password)
	if err != nil {
		return u, err
	}

	result, err := stmt.Exec(u.Email, password)
	if err != nil {

		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
				return u, ErrUserExists
			}
		}

		return u, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return u, err
	}
	u.ID = id

	return u, nil
}
