package service_test

import (
	"errors"
	"testing"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/pkg/scanner"
	"todo-cli-go/service"

	"github.com/stretchr/testify/assert"
)

func TestCategoryValidationService_IsUserCategory(t *testing.T) {
	testCases := []struct {
		name       string
		userID     uint
		categoryID uint
		expected   bool
		err        error
	}{
		{
			name:       "no category",
			categoryID: scanner.NoID,
			userID:     1,
			expected:   true,
		},
		{
			name:       "category owned by user",
			categoryID: 1,
			userID:     1,
			expected:   true,
		},
		{
			name:       "category not owned by user",
			categoryID: 1,
			userID:     2,
			expected:   false,
		},
		{
			name:       "category not found",
			categoryID: 123,
			userID:     1,
			err:        apperror.ErrCategoryNotFound,
		},
		{
			name:       "service have error",
			categoryID: 123,
			err:        serviceErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var conf bool
			if errors.Is(tc.err, serviceErr) {
				conf = true
			}
			categoryService := &MockCategoryMaster{error: conf}
			categoryValidationService := service.NewCategoryValidationService(categoryService)

			// 2. execution
			isOwned, err := categoryValidationService.IsUserCategory(tc.userID, tc.categoryID)

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				// error free test cases
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, isOwned)
			}
		})
	}
}

type MockCategoryMaster struct {
	error bool
}

func (m MockCategoryMaster) Get() ([]entity.Category, error) {
	if m.error {
		return nil, serviceErr
	}
	return categoryStorage, nil
}

func (m MockCategoryMaster) Create(title, color string) (*entity.Category, error) {
	panic("not implemented")
}

func (m MockCategoryMaster) Edit(id uint, title, color string) (*entity.Category, error) {
	panic("not implemented")
}
