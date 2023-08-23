package repository_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-cli-go/entity"
	apperror "todo-cli-go/error"
	"todo-cli-go/repository"
)

var categoryStorage = []entity.Category{
	{
		ID:     1,
		Title:  "Personal",
		Color:  "blue",
		UserID: 123,
	},
	{
		ID:     2,
		Title:  "Work",
		Color:  "Purple",
		UserID: 123,
	},
}

func TestCategoryRepository_GetByID(t *testing.T) {
	testCases := []struct {
		name       string
		categoryId uint
		err        error
		config     MockIOConfig
	}{
		{
			name:       "existing category",
			categoryId: 2,
			config:     MockIOConfig{read: true, write: true, delete: true},
		},
		{
			"not existing category",
			3,
			apperror.ErrCategoryNotFound,
			MockIOConfig{read: true, write: true, delete: true},
		},
		{
			"error on read",
			3,
			errors.New(ErrOnReading),
			MockIOConfig{read: false, write: true, delete: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler[entity.Category](&categoryStorage, tc.config)
			repo := repository.NewCategoryFileRepository(handler)

			// 2. execution
			category, err := repo.GetByID(tc.categoryId)

			// 3. assertion
			if tc.err != nil {
				assert.Equal(t, err, tc.err)

			} else {
				assert.NoError(t, err)
				assert.NotNil(t, category)
			}
		})
	}
}

func TestCategoryRepository_GetAll(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		storage  *[]entity.Category
		expected []entity.Category
		config   MockIOConfig
	}{
		{
			"have data in storage",
			nil,
			&categoryStorage,
			categoryStorage,
			MockIOConfig{read: true, write: true, delete: true},
		},
		{
			"empty data storage",
			nil,
			&[]entity.Category{},
			[]entity.Category{},
			MockIOConfig{read: true, write: true, delete: true},
		},
		{
			"error on reading",
			errors.New(ErrOnReading),
			&[]entity.Category{},
			nil,
			MockIOConfig{read: false, write: true, delete: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler[entity.Category](tc.storage, tc.config)
			repo := repository.NewCategoryFileRepository(handler)

			// 2. execution
			result, err := repo.GetAll()

			// 3. assertion
			if err != nil {
				assert.Equal(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result, "Expected and actual categories do not match")
			}
		})
	}
}

func TestCategoryRepository_Create(t *testing.T) {
	testCases := []struct {
		name string
		data struct {
			title  string
			color  string
			userId uint
		}
		err    error
		config MockIOConfig
	}{
		{
			"ordinary create",
			struct {
				title  string
				color  string
				userId uint
			}{
				title:  "New Category",
				color:  "Cyan",
				userId: 1,
			},
			nil,
			MockIOConfig{read: true, write: true, delete: true},
		},
		{
			"error on reading",
			struct {
				title  string
				color  string
				userId uint
			}{
				title:  "",
				color:  "",
				userId: 1,
			},
			errors.New(ErrOnReading),
			MockIOConfig{read: false, write: true, delete: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler[entity.Category](&categoryStorage, tc.config)
			repo := repository.NewCategoryFileRepository(handler)

			// 2. execution
			createdCategory, err := repo.Create(tc.data.title, tc.data.color, tc.data.userId)

			// 3. assertion
			if err != nil {
				assert.Equal(t, err, tc.err)
			} else {
				expected := entity.NewCategory(createdCategory.ID, tc.data.title, tc.data.color, tc.data.userId)

				assert.NoError(t, err)
				assert.Equal(t, expected, createdCategory)
			}

		})

	}
}

func TestCategoryRepository_Edit(t *testing.T) {
	type updateCategoryRequest struct {
		id    uint
		title string
		color string
	}

	testCases := []struct {
		name   string
		data   updateCategoryRequest
		config MockIOConfig
		err    error
	}{
		{
			"valid category",
			updateCategoryRequest{1, "updated", "cyan"},
			MockIOConfig{read: true, write: true, delete: true},
			nil,
		},
		{
			"not available category",
			updateCategoryRequest{id: 5},
			MockIOConfig{read: true, write: true, delete: true},
			apperror.ErrCategoryNotFoundToEdit,
		},
		{
			"error on read",
			updateCategoryRequest{},
			MockIOConfig{read: false, write: true, delete: true},
			errors.New(ErrOnReading),
		},
		{
			"error on delete or write",
			updateCategoryRequest{},
			MockIOConfig{read: true, write: true, delete: false},
			errors.New(ErrOnWritingOrReading),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler[entity.Category](&categoryStorage, tc.config)
			repo := repository.NewCategoryFileRepository(handler)

			// 2. execution
			editedCategory, err := repo.Edit(tc.data.id, tc.data.title, tc.data.color)

			// 3. assertion
			if err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				// getting the updated category from storage
				expected, err := repo.GetByID(tc.data.id)
				if err != nil {
					t.Fatalf("error on getting id, %s", err)
				}

				assert.NoError(t, err)
				assert.Equal(t, expected.Title, editedCategory.Title, "Edited title does not match")
				assert.Equal(t, expected.Color, editedCategory.Color, "Edited title does not match")
			}

		})
	}
}
