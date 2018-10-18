package model

type userStorage interface {
	getUserById(string) (*User, error)
	CreateUser(User) (*User, error)	
}

type User struct {
	ID string 
	Name string 
	Email	string
	Password	string
}