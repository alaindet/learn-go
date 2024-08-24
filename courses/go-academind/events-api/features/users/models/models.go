package models

type UserModel struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type UserModelDisplay struct {
	ID    int64
	Email string
}

func (u UserModel) display() UserModelDisplay {
	return UserModelDisplay{
		ID:    u.ID,
		Email: u.Email,
	}
}
