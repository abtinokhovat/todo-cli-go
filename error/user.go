package apperror

import "errors"

var (
	ErrUserWrongPasswordOrEmail = errors.New("either your entered email or password is wrong")
	ErrUserNotFound             = errors.New("user not found")
)
