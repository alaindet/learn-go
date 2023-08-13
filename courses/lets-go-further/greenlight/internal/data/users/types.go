package users

// TODO: Abstract with a generic model interface?
type UserModelInterface interface {
	Insert(movie *User) error
	GetByEmail(email string) (*User, error)
	Update(user *User) error
}
