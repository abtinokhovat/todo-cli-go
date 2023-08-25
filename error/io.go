package apperror

import "errors"

var (
	ErrNotCorrectDigit = errors.New("error: the entered value is not a valid digit")
)
