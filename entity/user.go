package entity

import "fmt"

type User struct {
	ID       uint
	Email    string
	Password string
}

func NewUser(id uint, email, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

func (u User) String() string {
	return fmt.Sprintf("#%d-%s", u.ID, u.Email)
}
