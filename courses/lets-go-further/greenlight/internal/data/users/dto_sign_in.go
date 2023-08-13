package users

type SignInUserData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *SignInUserData) ToUser() (*User, error) {
	user := &User{
		Name:      d.Name,
		Email:     d.Email,
		Activated: false,
	}

	err := user.Password.Set(d.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
