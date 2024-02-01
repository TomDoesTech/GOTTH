package store

type User struct {
	Email    string
	Password string
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}
