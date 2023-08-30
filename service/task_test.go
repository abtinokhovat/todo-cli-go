package service_test

import (
	"errors"
	"slices"
	"testing"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/pkg/date"
	"todo-cli-go/pkg/scanner"
	"todo-cli-go/service"
	"todo-cli-go/test/util"

	"github.com/stretchr/testify/assert"
)

var (
	taskStorage = []entity.Task{
		{
			ID:         28,
			Title:      "Prepare Presentation",
			DueDate:    nil,
			Done:       false,
			CategoryID: 2,
			UserID:     2,
		},
		{
			ID:         1,
			Title:      "Prepare Presentation",
			DueDate:    util_test.GetDate(2024, 11, 26),
			Done:       false,
			CategoryID: 2,
			UserID:     2,
		},
		{
			ID:         2,
			Title:      "Complete Golang Assignment",
			DueDate:    nil,
			Done:       true,
			CategoryID: 1,
			UserID:     1,
		},
		{
			ID:         4,
			Title:      "testable",
			DueDate:    util_test.GetDate(2024, 11, 27),
			Done:       true,
			CategoryID: 0,
			UserID:     2,
		},
		{
			ID:         5,
			Title:      "shopping",
			DueDate:    util_test.GetDate(2024, 11, 27),
			Done:       true,
			CategoryID: 1,
			UserID:     2,
		},
		{
			ID:         6,
			Title:      "fighting",
			DueDate:    date.Now(),
			Done:       false,
			CategoryID: 1,
			UserID:     2,
		},
		{
			ID:         7,
			Title:      "fighting-2",
			DueDate:    date.Now(),
			Done:       false,
			CategoryID: 1,
			UserID:     2,
		},
	}
	serviceErr = errors.New("service err")
)

func TestBuildTaskService(t *testing.T) {

}

func TestTaskService_Create(t *testing.T) {
	testCases := []struct {
		name   string
		task   entity.Task
		config mockCategoryValidatorConfig
		err    error
	}{
		{
			name:   "ordinary create with category and due date",
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: true},
			task: entity.Task{
				Title:      "",
				DueDate:    util_test.GetDate(2023, 11, 11),
				CategoryID: 1,
			},
		},
		{
			name:   "ordinary create with due date",
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: true},
			task: entity.Task{
				Title:   "",
				DueDate: util_test.GetDate(2023, 11, 11),
			},
		},
		{
			name:   "ordinary create with category",
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: true},
			task: entity.Task{
				CategoryID: 1,
			},
		},
		{
			name:   "not owned category",
			config: mockCategoryValidatorConfig{accept: false, serviceErr: false, found: true},
			task: entity.Task{
				CategoryID: 1,
			},
			err: apperror.ErrUnauthorized,
		},
		{
			name:   "category not found to create with",
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: false},
			task: entity.Task{
				CategoryID: 144,
			},
			err: apperror.ErrCategoryNotFound,
		},
		{
			name:   "error in validator service",
			config: mockCategoryValidatorConfig{accept: true, serviceErr: true, found: true},
			err:    serviceErr,
		},
		{
			name:   "repo error",
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: true},
			task:   entity.Task{},
			err:    repoErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var haveError bool
			if errors.Is(tc.err, repoErr) {
				haveError = true
			}

			repo := NewMockTaskRepository(haveError)
			validator := NewMockCategoryValidator(tc.config)
			srv := service.NewTaskService(&entity.User{ID: 2}, validator, repo)

			// 2. execution
			task, err := srv.Create(tc.task.Title, tc.task.DueDate, tc.task.CategoryID)

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				// error free test cases
				assert.NoError(t, err)
				assert.True(t, slices.Contains(taskStorage, *task))
			}
		})
	}
}

func TestTaskService_Edit(t *testing.T) {
	testCases := []struct {
		name   string
		id     uint
		task   service.TaskUpdate
		config mockCategoryValidatorConfig
		err    error
	}{
		{
			name: "ordinary edit",
			id:   1,
			task: service.TaskUpdate{
				Title:      pointer("Updated"),
				DueDate:    util_test.GetDate(2025, 12, 2),
				Done:       pointer(true),
				CategoryID: pointer[uint](12),
			},
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: true},
		},
		{
			name: "just edit title",
			id:   28,
			task: service.TaskUpdate{
				Title: pointer("Updated"),
			},
		},
		{
			name: "just edit due date",
			id:   1,
			task: service.TaskUpdate{
				DueDate: util_test.GetDate(2025, 12, 2),
			},
		},
		{
			name: "just edit done",
			id:   1,
			task: service.TaskUpdate{
				Done: pointer(true),
			},
		},
		{
			name: "remove category id",
			id:   5,
			task: service.TaskUpdate{
				CategoryID: pointer[uint](0),
			},
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: true},
		},
		{
			name: "just edit category id",
			id:   1,
			task: service.TaskUpdate{
				CategoryID: pointer[uint](12),
			},
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: true},
		},
		{
			name: "edit category id which is not owned by user",
			id:   1,
			task: service.TaskUpdate{
				CategoryID: pointer[uint](32),
			},
			config: mockCategoryValidatorConfig{accept: false, serviceErr: false, found: true},
			err:    apperror.ErrUnauthorized,
		},
		{
			name: "edit category id which will cause service err",
			id:   1,
			task: service.TaskUpdate{
				CategoryID: pointer[uint](32),
			},
			config: mockCategoryValidatorConfig{accept: true, serviceErr: true, found: true},
			err:    serviceErr,
		},
		{
			name: "edit category id which will do not exists",
			id:   1,
			task: service.TaskUpdate{
				CategoryID: pointer[uint](144),
			},
			config: mockCategoryValidatorConfig{accept: true, serviceErr: false, found: false},
			err:    apperror.ErrCategoryNotFound,
		},
		{
			name: "not available task edit",
			id:   120,
			err:  apperror.ErrTaskNotFoundToEdit,
		},
		{
			name: "not authorized to edit task",
			id:   2,
			err:  apperror.ErrUnauthorized,
		},
		{
			name: "repo error",
			id:   1,
			err:  repoErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var haveError bool
			if errors.Is(tc.err, repoErr) {
				haveError = true
			}

			repo := NewMockTaskRepository(haveError)
			validator := NewMockCategoryValidator(tc.config)
			srv := service.NewTaskService(&entity.User{ID: 2}, validator, repo)

			// 2. execution
			beforeEdit, _ := srv.GetByID(tc.id)
			editedTask, err := srv.Edit(tc.id, tc.task)

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				// check if the value was in update query and if it was not check the previous value
				if tc.task.Title == nil {
					assert.Equal(t, beforeEdit.Title, editedTask.Title)
				} else {
					assert.Equal(t, *tc.task.Title, editedTask.Title)
				}

				if tc.task.Done == nil {
					assert.Equal(t, beforeEdit.Done, editedTask.Done)
				} else {
					assert.Equal(t, *tc.task.Done, editedTask.Done)
				}

				if tc.task.DueDate == nil {
					assert.Equal(t, beforeEdit.DueDate, editedTask.DueDate)
				} else {
					assert.Equal(t, *tc.task.DueDate, *editedTask.DueDate)
				}

				if tc.task.CategoryID == nil {
					assert.Equal(t, beforeEdit.CategoryID, editedTask.CategoryID)
				} else {
					assert.Equal(t, *tc.task.CategoryID, editedTask.CategoryID)
				}
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	testCases := []struct {
		name   string
		userID uint
		err    error
	}{
		{
			name:   "ordinary get",
			userID: 2,
		},
		{
			name: "repo error",
			err:  repoErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var haveError bool
			if errors.Is(tc.err, repoErr) {
				haveError = true
			}

			repo := NewMockTaskRepository(haveError)
			srv := service.NewTaskService(&entity.User{ID: tc.userID}, nil, repo)

			// 2. execution
			tasks, err := srv.Get()

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				// error free test cases
				assert.NoError(t, err)
				for _, task := range tasks {
					assert.Equal(t, tc.userID, task.UserID)
				}
			}
		})
	}
}

func TestTaskService_GetByDate(t *testing.T) {
	testCases := []struct {
		name string
		date date.Date
		len  uint
		err  error
	}{
		{
			name: "ordinary get",
			date: *util_test.GetDate(2024, 11, 26),
			len:  2,
		},
		{
			name: "not available date",
			date: *util_test.GetDate(2110, 11, 26),
			len:  0,
		},
		{
			name: "repo error",
			err:  repoErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var haveError bool
			if errors.Is(tc.err, repoErr) {
				haveError = true
			}

			userID := uint(2)

			repo := NewMockTaskRepository(haveError)
			srv := service.NewTaskService(&entity.User{ID: userID}, nil, repo)

			// 2. execution
			tasks, err := srv.GetByDate(tc.date)

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				// error free test cases
				assert.NoError(t, err)
				for _, task := range tasks {
					assert.Equal(t, task.UserID, userID)
					assert.Equal(t, task.DueDate, tc.date)
				}
			}
		})
	}
}

func TestTaskService_GetToday(t *testing.T) {
	testCases := []struct {
		name string
		len  uint
		err  error
	}{
		{
			name: "ordinary get",
			len:  2,
		},
		{
			name: "repo error",
			err:  repoErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var haveError bool
			if errors.Is(tc.err, repoErr) {
				haveError = true
			}

			userID := uint(2)

			repo := NewMockTaskRepository(haveError)
			srv := service.NewTaskService(&entity.User{ID: userID}, nil, repo)

			// 2. execution
			tasks, err := srv.GetTodayTasks()

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				// error free test cases
				assert.NoError(t, err)
				for _, task := range tasks {
					assert.Equal(t, task.UserID, userID)
					assert.Equal(t, task.DueDate, date.Now())
				}
			}
		})
	}

}

func TestTaskService_Toggle(t *testing.T) {
	testCases := []struct {
		name     string
		taskID   uint
		user     entity.User
		expected bool
		err      error
	}{
		{
			name:     "ordinary toggle",
			taskID:   28,
			expected: true,
			user:     entity.User{ID: 2},
		},
		{
			name:     "ordinary toggle 2",
			taskID:   4,
			expected: false,
			user:     entity.User{ID: 2},
		},
		{
			name:   "not available task",
			taskID: 213,
			user:   entity.User{ID: 2},
			err:    apperror.ErrTaskNotFound,
		},
		{
			name:   "not user's task",
			taskID: 2,
			user:   entity.User{ID: 2},
			err:    apperror.ErrUnauthorized,
		},
		{
			name: "repo error",
			err:  repoErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			var haveError bool
			if errors.Is(tc.err, repoErr) {
				haveError = true
			}

			repo := NewMockTaskRepository(haveError)
			srv := service.NewTaskService(&entity.User{ID: tc.user.ID}, nil, repo)

			// 2. execution
			err := srv.Toggle(tc.taskID)

			// 3. assertion
			if tc.err != nil {
				// check for errors
				assert.Equal(t, tc.err, err)
			} else {
				afterToggle, err := repo.GetByID(tc.taskID)
				afterStatus := afterToggle.Done

				// error free test cases
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, afterStatus)
			}
		})
	}
}

func pointer[T any](value T) *T {
	return &value
}

type MockTaskRepository struct {
	haveError bool
}

func NewMockTaskRepository(haveError bool) *MockTaskRepository {
	return &MockTaskRepository{haveError: haveError}
}

func (r *MockTaskRepository) Create(title string, date *date.Date, categoryID, userID uint) (*entity.Task, error) {
	if r.haveError {
		return nil, repoErr
	}

	task := entity.NewTask(uint(len(taskStorage)+1), title, false, date, categoryID, userID)
	taskStorage = append(taskStorage, *task)

	return task, nil
}
func (r *MockTaskRepository) Edit(id uint, title string, done bool, date *date.Date, categoryID uint) (*entity.Task, error) {
	if r.haveError {
		return nil, repoErr
	}

	for i := 0; i < len(taskStorage); i++ {
		if taskStorage[i].ID == id {
			taskStorage[i].Title = title
			taskStorage[i].Done = done
			taskStorage[i].DueDate = date
			taskStorage[i].CategoryID = categoryID

			return &taskStorage[i], nil
		}

	}

	return nil, apperror.ErrTaskNotFoundToEdit
}
func (r *MockTaskRepository) GetByID(id uint) (*entity.Task, error) {
	if r.haveError {
		return nil, repoErr
	}

	for _, task := range taskStorage {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, apperror.ErrTaskNotFound
}
func (r *MockTaskRepository) GetAll() ([]entity.Task, error) {
	if r.haveError {
		return nil, repoErr
	}
	return taskStorage, nil
}

type mockCategoryValidator struct {
	config mockCategoryValidatorConfig
}

type mockCategoryValidatorConfig struct {
	// accept is true when category is owned by user
	accept bool
	// serviceErr is true to mock the service err
	serviceErr bool
	// if found is false the mocker will return not found err
	found bool
}

func NewMockCategoryValidator(config mockCategoryValidatorConfig) service.CategoryValidator {
	return &mockCategoryValidator{config: config}
}

func (m mockCategoryValidator) IsUserCategory(userID, categoryID uint) (bool, error) {
	if m.config.serviceErr {
		return false, serviceErr
	}
	if categoryID == scanner.NoID {
		return true, nil
	}
	if !m.config.found {
		return false, apperror.ErrCategoryNotFound
	}
	if m.config.accept {
		return true, nil
	}
	return false, apperror.ErrUnauthorized
}
