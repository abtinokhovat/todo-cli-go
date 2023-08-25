package apperror

import "errors"

var (
	ErrCategoryNotFound       = errors.New("error: category not found")
	ErrCategoryNotFoundToEdit = errors.New("error: category not found to edit")
)
