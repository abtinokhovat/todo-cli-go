package repository

import (
	"sync"
	"todo-cli-go/entity"
	apperror "todo-cli-go/error"

	fileHandler "github.com/abtinokhovat/file-handler-go"
)

const userStoragePath = "storage/user.json"

var (
	onceUser               sync.Once
	userRepositoryInstance *UserRepository
)

type IDGenerator interface {
	newID() uint
}

type UserStorageAdapter interface {
	Create(email, password string) (*entity.User, error)
	Get(email string) (*entity.User, error)
}

type UserRepository struct {
	handler fileHandler.FileIOHandler[entity.User]
}

func NewUserRepository(handler fileHandler.FileIOHandler[entity.User]) *UserRepository {
	return &UserRepository{handler: handler}
}
func GetUserRepository() *UserRepository {
	onceUser.Do(func() {
		serializer := fileHandler.NewJsonSerializer[entity.User]()
		handler := fileHandler.NewJsonIOHandler[entity.User](userStoragePath, serializer)
		userRepositoryInstance = NewUserRepository(handler)
	})
	return userRepositoryInstance
}

func (r *UserRepository) newID() uint {
	users, _ := r.handler.Read()
	return uint(len(users) + 1)
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

	return nil, apperror.ErrUserNotFound
}
