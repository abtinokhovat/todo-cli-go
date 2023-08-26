package apperror

import "errors"

var (
	ErrNotCorrectDigit    = errors.New("error: the entered value is not a valid digit")
	ErrOnReading          = errors.New("error on reading file")
	ErrOnWriting          = errors.New("error on writing file")
	ErrOnWritingOrReading = errors.New("error on deleting and writing file file")
)
