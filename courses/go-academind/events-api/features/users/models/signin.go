package models

import (
	"app/common/utils"
	"app/core/db"
)

var fetchUserPasswordSql = `
	SELECT "password"
	FROM "users"
	WHERE "email" = ?
`

func (u UserModel) ValidateCredentials() error {

	stmt, err := db.DB.Prepare(fetchUserPasswordSql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var hashedPassword string
	row := stmt.QueryRow(u.Email)
	err = row.Scan(&hashedPassword)
	if err != nil {
		return ErrInvalidCredentials
	}

	if !utils.CheckPasswordHash(hashedPassword, u.Password) {
		return ErrInvalidCredentials
	}

	return nil
}
