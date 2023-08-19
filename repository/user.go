package repository

import (
	"errors"
	"sync"
	"todo-cli-go/entity"

	h "github.com/abtinokhovat/file-handler-go"
)

const path = ""

var (
	once            sync.Once
	instance        *UserRepository
	ErrUserNotFound = errors.New("user not found")
)

type UserStorageAdapter interface {
	Create(email, password string) error
	Get(email string) (*entity.User, error)
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

func (r *UserRepository) Create(email, password string) error {
	// TODO: implement get new id method
	id := 0
	user := entity.NewUser(id, email, password)
	err := r.handler.WriteOne(*user)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) Get(email string) (*entity.User, error) {
	users, err := r.handler.Read()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, ErrUserNotFound
}
