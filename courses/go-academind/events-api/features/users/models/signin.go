package models

import (
	"app/common/jwt"
	"app/common/utils"
	"app/core/db"
	"time"
)

var fetchUserPasswordSql = `
	SELECT "id", "password"
	FROM "users"
	WHERE "email" = ?
`

func (u *UserModel) ValidateCredentials() (string, error) {

	stmt, err := db.DB.Prepare(fetchUserPasswordSql)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var userId int64
	var hashedPassword string
	row := stmt.QueryRow(u.Email)
	err = row.Scan(&userId, &hashedPassword)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !utils.CheckPasswordHash(hashedPassword, u.Password) {
		return "", ErrInvalidCredentials
	}

	jwt, err := jwt.Generate(u.Email, userId, 12*time.Hour)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
