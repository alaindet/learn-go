package models

import (
	"app/common/utils"
	"app/core/db"
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
		return u, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return u, err
	}
	u.ID = id

	return u, nil
}
