package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"todo-cli-go/entity"
	"todo-cli-go/repository"
)

var (
	ErrUserWrongPasswordOrEmail = errors.New("either your entered email or password is wrong")
)

type AuthService interface {
	Login(email, password string) (*entity.User, error)
	Register(email, password string) (*entity.User, error)
}

type UserService struct {
	repository repository.UserStorageAdapter
}

func BuildUserService() *UserService {
	repo := repository.GetUserRepository()
	return NewUserService(repo)
}

func NewUserService(repository repository.UserStorageAdapter) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) hashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}

func (s *UserService) Login(email, password string) (*entity.User, error) {
	// get user from repository
	user, err := s.repository.Get(email)
	if err != nil {
		return nil, err
	}

	// hash password and validate
	if user.Password == s.hashPassword(password) {
		return user, nil
	} else {
		return nil, ErrUserWrongPasswordOrEmail
	}
}

func (s *UserService) Register(email, password string) (*entity.User, error) {
	// hash password and create the user with repository
	hashed := s.hashPassword(password)
	createdUser, err := s.repository.Create(email, hashed)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
