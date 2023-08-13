package users

import (
	"database/sql"
	"errors"
	"greenlight/internal/data/common"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(user *User) error {

	stmt := `
		INSERT INTO
			users (name, email, password_hash, activated)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id, created_at, version
	`

	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Activated,
	}

	ctx, cancel := common.NewDatabaseContext()
	defer cancel()

	err := m.DB.QueryRowContext(ctx, stmt, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Version,
	)

	if isDuplicateEmailErr(err) {
		return common.ErrDuplicateEmail
	}

	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) GetByEmail(email string) (*User, error) {

	stmt := `
		SELECT
			id,
			created_at,
			name,
			email,
			password_hash,
			activated,
			version
		FROM
			users
		WHERE
			email = $1
	`

	var user User

	ctx, cancel := common.NewDatabaseContext()
	defer cancel()

	err := m.DB.QueryRowContext(ctx, stmt, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.Version,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, common.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) Update(user *User) error {

	stmt := `
		UPDATE
			users
		SET
			name = $1,
			email = $2,
			password_hash = $3,
			activated = $4,
			version = version + 1
		WHERE
			id = $5 AND version = $6
		RETURNING
			version
	`

	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Activated,
		user.ID,
		user.Version,
	}

	ctx, cancel := common.NewDatabaseContext()
	defer cancel()

	err := m.DB.QueryRowContext(ctx, stmt, args...).Scan(&user.Version)

	if isDuplicateEmailErr(err) {
		return common.ErrDuplicateEmail
	}

	if errors.Is(err, sql.ErrNoRows) {
		return common.ErrEditConflict
	}

	if err != nil {
		return err
	}

	return nil
}
