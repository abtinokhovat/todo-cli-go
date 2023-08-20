package entity

import "fmt"

type User struct {
	Id       uint
	Email    string
	Password string
}

func NewUser(id int, email, password string) *User {
	hashed := hash(password)
	return &User{
		Id:       uint(id),
		Email:    email,
		Password: hashed,
	}
}

func (u User) String() string {
	return fmt.Sprintf("#%d-%s", u.Id, u.Email)
}

func hash(str string) string {
	// TODO: implement hashing
	return str
}
