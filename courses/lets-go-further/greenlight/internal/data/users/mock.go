package users

type MockUserModel struct{}

func (m *MockUserModel) Insert(user *User) error {
	return nil
}

func (m *MockUserModel) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (m *MockUserModel) Update(movie *User) error {
	return nil
}
