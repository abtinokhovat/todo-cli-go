package apperror

import "errors"

var (
	ErrUserWrongPasswordOrEmail = errors.New("error: either your entered email or password is wrong")
	ErrUserNotFound             = errors.New("error: user not found")
	ErrUnauthorized             = errors.New("error: the request you are making is not authorized")
)
