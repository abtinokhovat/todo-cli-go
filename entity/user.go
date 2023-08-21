package entity

import "fmt"

type User struct {
	Id       uint
	Email    string
	Password string
}

func NewUser(id uint, email, password string) *User {
	return &User{
		Id:       id,
		Email:    email,
		Password: password,
	}
}

func (u User) String() string {
	return fmt.Sprintf("#%d-%s", u.Id, u.Email)
}
