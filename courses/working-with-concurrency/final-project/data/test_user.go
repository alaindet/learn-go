package data

import "time"

type TestUser struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

func testGetDummyUser() *User {
	return &User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "admin@example.com",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *TestUser) GetAll() ([]*User, error) {
	users := []*User{
		testGetDummyUser(),
	}
	return users, nil
}

func (u *TestUser) GetByEmail(email string) (*User, error) {
	return testGetDummyUser(), nil
}

func (u *TestUser) GetOne(id int) (*User, error) {
	return testGetDummyUser(), nil
}

func (u *TestUser) Update(user User) error {
	return nil
}

func (u *TestUser) Delete() error {
	return nil
}

func (u *TestUser) DeleteByID(id int) error {
	return nil
}

func (u *TestUser) Insert(user User) (int, error) {
	newID := 2
	return newID, nil
}

func (u *TestUser) ResetPassword(password string) error {
	return nil
}

func (u *TestUser) PasswordMatches(plainText string) (bool, error) {
	return true, nil
}
