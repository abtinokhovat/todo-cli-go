package service

import (
	"errors"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/pkg/date"
	"todo-cli-go/repository"
)

type TaskUpdate struct {
	Title      *string
	Done       *bool
	DueDate    *date.Date
	CategoryID *uint
}

type TaskMaster interface {
	Create(title string, dueDate *date.Date, categoryID uint) (*entity.Task, error)
	Edit(id uint, update TaskUpdate) (*entity.Task, error)
	Get() ([]entity.Task, error)
	GetByID(id uint) (*entity.Task, error)
	GetTodayTasks() ([]entity.Task, error)
	GetByDate(date date.Date) ([]entity.Task, error)
	Toggle(id uint) error
}

type TaskService struct {
	user              *entity.User
	repo              repository.TaskStorageAdapter
	validationService CategoryValidator
}

func BuildTaskService(user *entity.User, validation CategoryValidator) *TaskService {
	repo := repository.GetTaskFileRepository()
	return NewTaskService(user, validation, repo)
}

func NewTaskService(user *entity.User, validation CategoryValidator, repository repository.TaskStorageAdapter) *TaskService {
	return &TaskService{
		user:              user,
		repo:              repository,
		validationService: validation,
	}
}

func (s *TaskService) Create(title string, dueDate *date.Date, categoryID uint) (*entity.Task, error) {
	isOwned, err := s.validationService.IsUserCategory(s.user.ID, categoryID)
	if !isOwned {
		return nil, err
	}
	task, err := s.repo.Create(title, dueDate, categoryID, s.user.ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (s *TaskService) Edit(id uint, update TaskUpdate) (*entity.Task, error) {
	task, err := s.repo.GetByID(id)
	if errors.Is(err, apperror.ErrTaskNotFound) {
		return nil, apperror.ErrTaskNotFoundToEdit
	} else if err != nil {
		return nil, err
	}

	if task.UserID != s.user.ID {
		return nil, apperror.ErrUnauthorized
	}

	// don't update fields if they were with zero values
	if update.Done != nil {
		task.Done = *update.Done
	}
	if update.Title != nil {
		task.Title = *update.Title
	}
	if update.DueDate != nil {
		task.DueDate = update.DueDate
	}
	if update.CategoryID != nil {
		// check if category belongs to user
		isOwned, err := s.validationService.IsUserCategory(s.user.ID, *update.CategoryID)
		if err != nil {
			return nil, err
		}
		if !isOwned {
			return nil, apperror.ErrUnauthorized
		}
		task.CategoryID = *update.CategoryID
	}

	edited, err := s.repo.Edit(id, task.Title, task.Done, task.DueDate, task.CategoryID)
	if err != nil {
		return nil, err
	}

	return edited, nil
}
func (s *TaskService) Get() ([]entity.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var userTasks []entity.Task
	for _, task := range tasks {
		if task.UserID == s.user.ID {
			userTasks = append(userTasks, task)
		}
	}

	return userTasks, nil
}
func (s *TaskService) GetByID(id uint) (*entity.Task, error) {
	tasks, err := s.Get()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, apperror.ErrTaskNotFound
}
func (s *TaskService) GetTodayTasks() ([]entity.Task, error) {
	today := date.Now()
	todayTasks, err := s.GetByDate(*today)

	if err != nil {
		return nil, err
	}

	return todayTasks, nil
}
func (s *TaskService) GetByDate(date date.Date) ([]entity.Task, error) {
	tasks, err := s.Get()
	if err != nil {
		return nil, err
	}

	var selectedDateTasks []entity.Task
	for _, task := range tasks {
		if task.DueDate != nil && task.DueDate.IsSameDate(date) {
			selectedDateTasks = append(selectedDateTasks, task)
		}
	}

	return selectedDateTasks, nil
}
func (s *TaskService) Toggle(id uint) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// return error if task was not the user's task
	if task.UserID != s.user.ID {
		return apperror.ErrUnauthorized
	}

	var isDone bool
	if !task.Done {
		isDone = true
	}

	updateReq := TaskUpdate{
		Done: &isDone,
	}

	// edit the task with toggled isDone
	_, err = s.Edit(id, updateReq)

	return err
}
