package service_test

import (
	"errors"
	"slices"
	"testing"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/service"

	"github.com/stretchr/testify/assert"
)

var (
	errRepo         = errors.New("repository error")
	categoryStorage = []entity.Category{
		{
			Title:  "personal",
			Color:  "red",
			UserID: 1,
		},
		{
			ID:     1,
			Title:  "work",
			Color:  "purple",
			UserID: 1,
		},
		{
			ID:     3,
			Title:  "life",
			Color:  "cyan",
			UserID: 1,
		},
		{
			ID:     2,
			Title:  "personal",
			Color:  "red",
			UserID: 2,
		},
	}
)

func TestCategoryService_Create(t *testing.T) {
	testCases := []struct {
		name  string
		user  entity.User
		title string
		color string
		err   error
	}{
		{
			name:  "ordinary creation",
			user:  entity.User{ID: 1},
			title: "new_category",
			color: "new_color",
		},
		{
			name: "creation with repo error",
			user: entity.User{ID: 1},
			err:  errRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var haveError bool
			if tc.err != nil {
				haveError = true
			}
			repo := NewMockCategoryRepository(haveError)
			srv := service.NewCategoryService(&tc.user, repo)

			// 2. execution
			created, err := srv.Create(tc.title, tc.color)

			// 3. assertion
			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, slices.Contains(categoryStorage, *created))
			}
		})
	}

}

func TestCategoryService_Edit(t *testing.T) {
	testCases := []struct {
		name      string
		category  entity.Category
		err       error
		repoError bool
	}{
		{
			name:     "ordinary edit",
			category: entity.Category{ID: 1, Title: "updated category", Color: "updated color", UserID: 1},
		},
		{
			name:     "ordinary edit 2",
			category: entity.Category{ID: 3, Title: "updated category", Color: "updated color", UserID: 1},
		},
		{
			name:     "not found category",
			category: entity.Category{ID: 5, UserID: 1},
			err:      apperror.ErrCategoryNotFound,
		},
		{
			name:     "category not belong to user to edit",
			category: entity.Category{ID: 2, UserID: 1},
			err:      apperror.ErrUnauthorized,
		},
		{
			name:      "repository error",
			category:  entity.Category{ID: 1, UserID: 1},
			err:       errRepo,
			repoError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := NewMockCategoryRepository(tc.repoError)
			srv := service.NewCategoryService(&entity.User{ID: tc.category.UserID}, repo)

			// 2. execution
			edited, err := srv.Edit(tc.category.ID, tc.category.Title, tc.category.Color)

			// 3. assertion
			if err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)

				// checking properties for equality
				assert.Equal(t, tc.category.Title, edited.Title)
				assert.Equal(t, tc.category.Color, edited.Color)
			}

		})
	}
}

func TestCategoryService_Get(t *testing.T) {
	testCases := []struct {
		name      string
		userID    uint
		err       error
		repoError bool
	}{
		{
			name:   "ordinary get",
			userID: 1,
		},
		{
			name:      "repo error",
			userID:    1,
			err:       errRepo,
			repoError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := NewMockCategoryRepository(tc.repoError)
			srv := service.NewCategoryService(&entity.User{ID: tc.userID}, repo)

			// 2. execution
			categories, err := srv.Get()

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)

			} else {
				// error free test cases
				assert.NoError(t, err)
				for _, category := range categories {
					assert.Equal(t, tc.userID, category.UserID)
				}
			}
		})
	}
}

type MockCategoryRepository struct {
	haveError bool
}

func NewMockCategoryRepository(haveError bool) *MockCategoryRepository {
	return &MockCategoryRepository{haveError: haveError}
}

func (r *MockCategoryRepository) Create(title, color string, userID uint) (*entity.Category, error) {
	if r.haveError == true {
		return nil, errRepo
	}
	// make a category
	category := entity.NewCategory(uint(len(categoryStorage)+1), title, color, userID)

	// append it to storage
	categoryStorage = append(categoryStorage, *category)

	return category, nil
}
func (r *MockCategoryRepository) Edit(id uint, title, color string) (*entity.Category, error) {
	if r.haveError == true {
		return nil, errRepo
	}

	for i := 0; i < len(categoryStorage); i++ {
		if categoryStorage[i].ID == id {
			categoryStorage[i].Title = title
			categoryStorage[i].Color = color

			return &categoryStorage[i], nil
		}
	}

	return nil, apperror.ErrCategoryNotFoundToEdit
}
func (r *MockCategoryRepository) GetByID(id uint) (*entity.Category, error) {
	if r.haveError == true {
		return nil, errRepo
	}

	for _, category := range categoryStorage {
		if category.ID == id {
			return &category, nil
		}
	}

	return nil, apperror.ErrCategoryNotFound
}
func (r *MockCategoryRepository) GetAll() ([]entity.Category, error) {
	if r.haveError == true {
		return nil, errRepo
	}

	return categoryStorage, nil
}
