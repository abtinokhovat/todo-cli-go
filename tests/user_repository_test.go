package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-cli-go/entity"
	"todo-cli-go/repository"
)

var userStorage = []entity.User{
	{
		Email:    "user1@example.com",
		Password: "password123",
		Id:       1,
	},
	{
		Email:    "user2@example.com",
		Password: "securepass",
		Id:       2,
	},
}

func TestUserRepository_Create(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		mockHandler := &MockUserIOHandler{}
		userRepo := repository.NewUserRepository(mockHandler)

		email := "test@example.com"
		password := "password"

		// 2. execution
		err := userRepo.Create(email, password)

		// 3. assertion
		assert.NoError(t, err)
		assert.Equal(t, 3, len(userStorage), "user length expected to be 3 but was %d", len(userStorage))
	})

}

func TestUserRepository_Get(t *testing.T) {
	mockHandler := &MockUserIOHandler{}
	userRepo := repository.NewUserRepository(mockHandler)

	// Adding test data to the mock handler
	email := "test@example.com"
	password := "password"
	user := entity.NewUser(1, email, password)
	err := mockHandler.WriteOne(*user)
	if err != nil {
		return
	}

	// Test existing user
	foundUser, err := userRepo.Get(email)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, email, foundUser.Email)

	// Test non-existing user
	nonExistentEmail := "nonexistent@example.com"
	_, err = userRepo.Get(nonExistentEmail)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrUserNotFound))
}

// MockUserIOHandler is a mock implementation of the FileIOHandler interface
type MockUserIOHandler struct {
}

func (h *MockUserIOHandler) Read() ([]entity.User, error) {
	return userStorage, nil
}
func (h *MockUserIOHandler) WriteOne(data entity.User) error {
	userStorage = append(userStorage, data)
	return nil
}
func (h *MockUserIOHandler) DeleteAndWrite(data []entity.User) error {
	userStorage = data
	return nil
}
func (h *MockUserIOHandler) DeleteAll() error {
	userStorage = []entity.User{}
	return nil
}
