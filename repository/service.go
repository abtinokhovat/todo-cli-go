package repository

import "todo-cli-go/entity"

type UserStorageAdapter interface {
	Create(email, password string) error
	Get(email, password string) *entity.User
}
