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
	repository repository.UserStorageAdapter
}

func BuildUserService() *UserService {
	//TODO complete implementation
	repo := repository.UserRepository{}
	return NewUserService(&repo)
}

func NewUserService(repository repository.UserStorageAdapter) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Login(email, password string) (*entity.User, error) {
	//TODO implement me
	return &entity.User{
		Id:       0,
		Email:    email,
		Password: password,
	}, nil
}

func (u *UserService) Register(name, email string) error {
	//TODO implement me
	return nil
}

func hash(str string) string {
	//TODO implement me
	panic("implement me")
}
