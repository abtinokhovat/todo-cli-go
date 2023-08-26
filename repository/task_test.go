package repository_test

import (
	"errors"
	"testing"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/pkg/date"
	"todo-cli-go/repository"

	"github.com/stretchr/testify/assert"
)

var taskStorage = []entity.Task{
	{
		ID:         2,
		Title:      "Complete Golang Assignment",
		DueDate:    nil,
		Done:       true,
		CategoryID: 1,
		UserID:     1,
	},
	{
		ID:         1,
		Title:      "Prepare Presentation",
		DueDate:    getDate(2024, 11, 27),
		Done:       false,
		CategoryID: 2,
		UserID:     2,
	},
	{
		ID:         4,
		Title:      "testable",
		DueDate:    getDate(2024, 11, 27),
		Done:       false,
		CategoryID: 0,
		UserID:     2,
	},
}

func TestGetTodoFileRepository(t *testing.T) {
	t.Run("test singelton", func(t *testing.T) {
		// 1. setup
		var repoArray []*repository.TaskFileRepository

		// 2. execution
		for i := 0; i < 30; i++ {
			repo := repository.GetTaskFileRepository()
			repoArray = append(repoArray, repo)
		}

		// 3. assertion
		for i := 0; i < len(repoArray)-1; i++ {
			assert.Equal(t, repoArray[i], repoArray[i+1])
		}
	})
}

func TestTodoFileRepository_Create(t *testing.T) {
	testCases := []struct {
		name   string
		config MockIOConfig
		task   entity.Task
		err    error
	}{
		{
			name:   "ordinary creation with due date and category id",
			config: MockIOConfig{read: true, write: true, delete: true},
			task: entity.Task{
				Title:      "test-1",
				DueDate:    date.Now(),
				CategoryID: 1,
				UserID:     2,
			},
		},
		{
			name:   "ordinary creation without due date",
			config: MockIOConfig{read: true, write: true, delete: true},
			task: entity.Task{
				Title:      "test-1",
				CategoryID: 1,
				UserID:     2,
			},
		},
		{
			name:   "ordinary creation without category id",
			config: MockIOConfig{read: true, write: true, delete: true},
			task: entity.Task{
				Title:   "test-1",
				DueDate: date.Now(),
				UserID:  2,
			},
		},
		{
			name:   "error on read on new id",
			config: MockIOConfig{read: false, write: true, delete: true},
			err:    errors.New(ErrOnReading),
		},
		{
			name:   "error on write",
			config: MockIOConfig{read: true, write: false, delete: true},
			err:    errors.New(ErrOnWriting),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler(&taskStorage, tc.config)
			repo := repository.NewTaskFileRepository(handler)

			// 2. execution
			task, err := repo.Create(tc.task.Title, tc.task.DueDate, tc.task.CategoryID, tc.task.UserID)

			// 3. assertion
			if tc.err != nil {
				// check for errors
			} else {
				expected := entity.NewTask(task.ID, tc.task.Title, false, tc.task.DueDate, tc.task.CategoryID, tc.task.UserID)

				// error free test cases
				assert.NoError(t, err)
				assert.Equal(t, expected, task)
			}
		})
	}
}

func TestTodoFileRepository_Edit(t *testing.T) {
	testCases := []struct {
		name   string
		config MockIOConfig
		task   entity.Task
		err    error
	}{
		{
			name:   "ordinary edit",
			config: MockIOConfig{read: true, write: true, delete: true},
			task: entity.Task{
				ID:      4,
				Title:   "Updated",
				DueDate: getDate(2025, 12, 2),
				Done:    true,
			},
		},
		{
			name:   "not available task edit",
			config: MockIOConfig{read: true, write: true, delete: true},
			task:   entity.Task{ID: 7},
			err:    apperror.ErrTaskNotFoundToEdit,
		},
		{
			name:   "error on read",
			config: MockIOConfig{read: false, write: true, delete: true},
			err:    errors.New(ErrOnReading),
		},
		{
			name:   "not deleting and writing",
			config: MockIOConfig{read: true, write: false, delete: false},
			task:   entity.Task{ID: 4},
			err:    errors.New(ErrOnWritingOrReading),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler(&taskStorage, tc.config)
			repo := repository.NewTaskFileRepository(handler)

			// 2. execution
			task, err := repo.Edit(tc.task.ID, tc.task.Title, tc.task.Done, tc.task.DueDate, tc.task.CategoryID)

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				expected, err := repo.GetByID(tc.task.ID)

				// error free test cases
				assert.NoError(t, err)
				assert.Equal(t, expected, task)
			}
		})
	}
}

func TestTodoFileRepository_GetById(t *testing.T) {
	testCases := []struct {
		name   string
		taskID uint
		err    error
		config MockIOConfig
	}{
		{
			name:   "existing task",
			taskID: 2,
			config: MockIOConfig{read: true, write: true, delete: true},
		},
		{
			name:   "not existing category",
			taskID: 7,
			err:    apperror.ErrTaskNotFound,
			config: MockIOConfig{read: true, write: true, delete: true},
		},
		{
			name:   "error on read",
			taskID: 2,
			err:    errors.New(ErrOnReading),
			config: MockIOConfig{read: false, write: true, delete: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler(&taskStorage, tc.config)
			repo := repository.NewTaskFileRepository(handler)

			// 2. execution
			task, err := repo.GetByID(tc.taskID)

			// 3. assertion
			if tc.err != nil {
				assert.Equal(t, err, tc.err)

			} else {
				assert.NoError(t, err)
				assert.NotNil(t, task)
			}
		})
	}

}

func TestTodoFileRepository_GetAll(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		storage  *[]entity.Task
		expected []entity.Task
		config   MockIOConfig
	}{
		{
			"have data in storage",
			nil,
			&taskStorage,
			taskStorage,
			MockIOConfig{read: true, write: true, delete: true},
		},
		{
			"empty data storage",
			nil,
			&[]entity.Task{},
			[]entity.Task{},
			MockIOConfig{read: true, write: true, delete: true},
		},
		{
			"error on reading",
			errors.New(ErrOnReading),
			&[]entity.Task{},
			nil,
			MockIOConfig{read: false, write: true, delete: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			handler := NewMockIOHandler[entity.Task](tc.storage, tc.config)
			repo := repository.NewTaskFileRepository(handler)

			// 2. execution
			result, err := repo.GetAll()

			// 3. assertion
			if err != nil {
				assert.Equal(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result, "Expected and actual tasks do not match")
			}
		})
	}
}

func getDate(year, month, day uint) *date.Date {
	d, _ := date.NewDate(year, month, day)
	return d
}
