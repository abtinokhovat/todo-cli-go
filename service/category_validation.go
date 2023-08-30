package service

import (
	"todo-cli-go/error"
	"todo-cli-go/pkg/scanner"
)

type CategoryValidator interface {
	IsUserCategory(userID, categoryID uint) (bool, error)
}

type CategoryValidationService struct {
	categoryService *CategoryService
}

func NewCategoryValidationService(category *CategoryService) CategoryValidator {
	return &CategoryValidationService{
		categoryService: category,
	}
}

func (s *CategoryValidationService) IsUserCategory(userID, categoryID uint) (bool, error) {
	// if there was no id return true
	if categoryID == scanner.NoID {
		return true, nil
	}
	// Validation logic to check if the category belongs to the user
	categories, err := s.categoryService.Get()
	if err != nil {
		return false, err
	}

	for _, category := range categories {
		if category.ID == categoryID {
			return category.UserID == userID, nil
		}
	}

	return false, apperror.ErrCategoryNotFound
}
