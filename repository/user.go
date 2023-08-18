package repository

import (
	h "github.com/abtinokhovat/file-handler-go"
	"sync"
	"todo-cli-go/entity"
)

var (
	once     sync.Once
	instance *UserRepository
)

const path = ""

type UserStorageAdapter interface {
	Create(email, password string) error
	Get(email, password string) *entity.User
}

type UserRepository struct {
	handler h.FileIOHandler[entity.User]
}

func NewUserRepository(handler h.FileIOHandler[entity.User]) *UserRepository {
	return &UserRepository{handler: handler}
}

func GetUserRepository() *UserRepository {
	once.Do(func() {
		serializer := h.NewJsonSerializer[entity.User]()
		handler := h.NewJsonIOHandler[entity.User](path, serializer)
		instance = NewUserRepository(handler)
	})
	return instance
}

func (u *UserRepository) Create(email, password string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Get(email, password string) *entity.User {
	//TODO implement me
	panic("implement me")
}
