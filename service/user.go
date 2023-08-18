package service

import (
	"todo-cli-go/entity"
	"todo-cli-go/repository"
)

type AuthService interface {
	Login(email, password string) (*entity.User, error)
	Register(name, email string) error
}

type UserService struct {
	repository *repository.UserStorageAdapter
}

func (u *UserService) Login(email, password string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Register(name, email string) error {
	//TODO implement me
	panic("implement me")
}

func hash(str string) string {
	panic("implement me")
}
