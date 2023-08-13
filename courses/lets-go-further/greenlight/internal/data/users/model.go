package users

import (
	"database/sql"
	"greenlight/internal/data/common"
	"strings"
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

// TODO
func (m *UserModel) GetByEmail(email string) *User {
	return nil
}

// TODO
func (m *UserModel) Update(user *User) error {
	return nil
}

// TODO
func isDuplicateEmailErr(err error) bool {
	duplicateErrMessage := "duplicate key value" // <-- TODO
	return strings.Contains(err.Error(), duplicateErrMessage)
}
