package apperror

import "errors"

var (
	ErrTaskNotFoundToEdit = errors.New("error: no task found for editing")
	ErrTaskNotFound       = errors.New("error: task not found")
)
