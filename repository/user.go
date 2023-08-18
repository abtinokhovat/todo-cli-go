package repository

import "todo-cli-go/entity"

type UserStorageAdapter interface {
	Create(email, password string) error
	Get(email, password string) *entity.User
}

type UserRepository struct {
}

func (u *UserRepository) Create(email, password string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Get(email, password string) *entity.User {
	//TODO implement me
	panic("implement me")
}
