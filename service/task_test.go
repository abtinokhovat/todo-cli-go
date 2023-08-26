package service_test

import (
	"testing"
	"todo-cli-go/pkg/date"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/test/util"
)

var (
	taskStorage = []entity.Task{
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
			DueDate:    util_test.GetDate(2024, 11, 26),
			Done:       false,
			CategoryID: 2,
			UserID:     2,
		},
		{
			ID:         4,
			Title:      "testable",
			DueDate:    util_test.GetDate(2024, 11, 27),
			Done:       false,
			CategoryID: 0,
			UserID:     2,
		},
		{
			ID:         5,
			Title:      "shopping",
			DueDate:    util_test.GetDate(2024, 11, 27),
			Done:       true,
			CategoryID: 0,
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
)

func TestBuildTaskService(t *testing.T) {

}

func TestTaskService_Create(t *testing.T) {
	testCases := []struct {
		name string
		task entity.Task
		err  error
	}{
		{
			name: "ordinary create with category and due date",
			task: entity.Task{
				Title:      "",
				DueDate:    nil,
				CategoryID: 0,
			},
		},
		{
			name: "ordinary create with due date",
			task: entity.Task{
				Title:      "",
				DueDate:    nil,
				CategoryID: 0,
			},
		},
		{
			name: "ordinary create with category",
			task: entity.Task{
				Title:      "",
				DueDate:    nil,
				CategoryID: 0,
			},
		},
		{
			name: "repo error",
			task: entity.Task{},
			err:  errRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			// 2. execution
			// 3. assertion
			if tc.err != nil {
				// check for errors
			} else {
				// error free test cases
			}

			// TODO: implement me
			panic("implement me")
		})
	}
}

func TestTaskService_Edit(t *testing.T) {
	testCases := []struct {
		name string
		task entity.Task
		err  error
	}{
		{
			name: "ordinary edit",
			task: entity.Task{
				ID:      2,
				Title:   "Updated",
				DueDate: util_test.GetDate(2025, 12, 2),
				Done:    true,
			},
		},
		{
			name: "not available task edit",
			task: entity.Task{ID: 120},
			err:  apperror.ErrTaskNotFoundToEdit,
		},
		{
			name: "not authorized to edit task",
			task: entity.Task{ID: 2},
			err:  apperror.ErrUnauthorized,
		},
		{
			name: "repo error",
			task: entity.Task{ID: 120},
			err:  errRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			// 2. execution
			// 3. assertion
			if tc.err != nil {
				// check for errors
			} else {
				// error free test cases
			}

			// TODO: implement me
			panic("implement me")
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	testCases := []struct {
		name      string
		userID    uint
		err       error
		repoError bool
	}{
		{
			name:   "ordinary get",
			userID: 2,
		},
		{
			name:      "repo error",
			err:       errRepo,
			repoError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			// 2. execution
			// 3. assertion
			if tc.err != nil {
				// check for errors
			} else {
				// error free test cases
			}

			// TODO: implement me
			panic("implement me")
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
			err:  errRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			// 2. execution
			// 3. assertion
			if tc.err != nil {
				// check for errors
			} else {
				// error free test cases
			}

			// TODO: implement me
			panic("implement me")
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
			err:  errRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			// 2. execution
			// 3. assertion
			if tc.err != nil {
				// check for errors
			} else {
				// error free test cases
			}

			// TODO: implement me
			panic("implement me")
		})
	}

}

func TestTaskService_Toggle(t *testing.T) {
	testCases := []struct {
		name   string
		taskID uint
		user   entity.User
		err    error
	}{
		{
			name:   "ordinary toggle",
			taskID: 1,
			user:   entity.User{ID: 2},
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
			err:  errRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			// 2. execution
			// 3. assertion
			if tc.err != nil {
				// check for errors
			} else {
				// error free test cases
			}

			// TODO: implement me
			panic("implement me")
		})
	}
}

type MockTaskRepository struct {
	haveError bool
}

func (r *MockTaskRepository) Create(title string, date *date.Date, categoryID, userID uint) (*entity.Task, error) {
	if r.haveError {
		return nil, errRepo
	}

	task := entity.NewTask(uint(len(taskStorage)+1), title, false, date, categoryID, userID)
	taskStorage = append(taskStorage, *task)

	return task, nil
}
func (r *MockTaskRepository) Edit(id uint, title string, done bool, date *date.Date, categoryID uint) (*entity.Task, error) {
	if r.haveError {
		return nil, errRepo
	}

	for i := 0; i < len(taskStorage); i++ {
		if taskStorage[i].ID == id {
			taskStorage[i].Title = title
			taskStorage[i].Done = done
			taskStorage[i].DueDate = date
			taskStorage[i].CategoryID = categoryID
		}
	}

	return nil, apperror.ErrTaskNotFoundToEdit
}
func (r *MockTaskRepository) GetByID(id uint) (*entity.Task, error) {
	if r.haveError {
		return nil, errRepo
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
		return nil, errRepo
	}
	return taskStorage, nil
}
