package service_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
	"todo-cli-go/entity"
	apperror "todo-cli-go/error"
	"todo-cli-go/service"
)

var userStorage = []entity.User{
	{
		0,
		"abtin@new.com",
		"202cb962ac59075b964b07152d234b70",
	},
}

func TestRegister(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
	}{
		{
			"ordinary registration",
			"test@example.ir",
			"123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := MockUserRepository{}
			srv := service.NewUserService(repo)

			// 2. execution
			registeredUser, err := srv.Register(tc.email, tc.password)

			// 3. assertion
			assert.NoError(t, err)
			assert.True(t, slices.Contains(userStorage, *registeredUser))
		})
	}

}

func TestLogin(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
		err      error
	}{
		{
			"existing user",
			"abtin@new.com",
			"123",
			nil,
		},
		{
			"wrong password",
			"abtin@new.com",
			"234",
			apperror.ErrUserWrongPasswordOrEmail,
		},
		{
			"not existing user",
			"abtin@new.ir",
			"234",
			apperror.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := MockUserRepository{}
			srv := service.NewUserService(repo)

			// 2. execution
			user, err := srv.Login(tc.email, tc.password)
			if err != nil {
				return
			}

			// 3. assertion
			if tc.err == nil {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			} else {
				assert.Error(t, tc.err)
				assert.Nil(t, user)
			}
		})
	}
}

type MockUserRepository struct {
}

func (m MockUserRepository) Create(email, password string) (*entity.User, error) {
	// create a user
	user := entity.NewUser(uint(len(userStorage)+1), email, password)
	// add user to storage
	userStorage = append(userStorage, *user)
	return user, nil
}

func (m MockUserRepository) Get(email string) (*entity.User, error) {
	// search for user in storage
	for _, user := range userStorage {
		if user.Email == email {
			return &user, nil
		}
	}

	// if user was not found return error
	return nil, apperror.ErrUserNotFound
}
