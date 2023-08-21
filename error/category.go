package apperror

import "errors"

var (
	ErrCategoryNotFound       = errors.New("category not found")
	ErrCategoryNotFoundToEdit = errors.New("category not found to edit")
)
