package repository_test

import (
	"slices"
	"testing"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/repository"

	"github.com/stretchr/testify/assert"
)

var userStorage = []entity.User{
	{
		Email:    "user1@example.com",
		Password: "password123",
		ID:       1,
	},
	{
		Email:    "user2@example.com",
		Password: "securepass",
		ID:       2,
	},
}

func TestGetUserRepository(t *testing.T) {
	t.Run("test singelton", func(t *testing.T) {
		// 1. setup
		var repoArray []*repository.UserRepository

		// 2. execution
		for i := 0; i < 30; i++ {
			repo := repository.GetUserRepository()
			repoArray = append(repoArray, repo)
		}

		// 3. assertion
		for i := 0; i < len(repoArray)-1; i++ {
			assert.Equal(t, repoArray[i], repoArray[i+1])
		}
	})
}

func TestUserRepository_Create(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		handler := NewMockIOHandler(&userStorage, MockIOConfig{read: true, write: true, delete: true})
		userRepo := repository.NewUserRepository(handler)

		email := "test@example.com"
		password := "password"

		// 2. execution
		createdUser, err := userRepo.Create(email, password)

		// 3. assertion
		assert.NoError(t, err)
		assert.Equal(t, 3, len(userStorage), "user length expected to be 3 but was %d", len(userStorage))
		assert.True(t, slices.Contains(userStorage, *createdUser))
	})

}

func TestUserRepository_Get(t *testing.T) {

	testCases := []struct {
		name     string
		email    string
		expected string
		err      error
	}{
		{
			"available user",
			userStorage[1].Email,
			userStorage[1].String(),
			nil,
		},
		{
			"not existed user",
			"nonexistent@example.com",
			"",
			apperror.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler(&userStorage, MockIOConfig{read: true, write: true, delete: true})
			userRepo := repository.NewUserRepository(handler)

			// 2. execution
			foundUser, err := userRepo.Get(tc.email)

			// 3. assertion
			// on error
			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, foundUser)
				assert.Equal(t, tc.expected, foundUser.String())
			}
		})

	}
}
