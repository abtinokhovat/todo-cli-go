package repository

import (
	"errors"
	"sync"
	"todo-cli-go/entity"

	h "github.com/abtinokhovat/file-handler-go"
)

const path = "storage/user.json"

var (
	once            sync.Once
	instance        *UserRepository
	ErrUserNotFound = errors.New("user not found")
)

type UserStorageAdapter interface {
	Create(email, password string) (*entity.User, error)
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

func (r *UserRepository) newID() int {
	users, _ := r.handler.Read()
	return len(users) + 1
}
func (r *UserRepository) Create(email, password string) (*entity.User, error) {
	// get new ID
	id := r.newID()

	// make user
	user := entity.NewUser(id, email, password)

	// write user with file handler
	err := r.handler.WriteOne(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
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
