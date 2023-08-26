package service

import (
	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/pkg/date"
	"todo-cli-go/repository"
)

type TaskService struct {
	user *entity.User
	repo repository.TaskStorageAdapter
}

func BuildTaskService(user *entity.User) *TaskService {
	repo := repository.GetTaskFileRepository()
	return NewTaskService(user, repo)
}

func NewTaskService(user *entity.User, repository repository.TaskStorageAdapter) *TaskService {
	return &TaskService{
		user: user,
		repo: repository,
	}
}

func (s *TaskService) Create(title string, dueDate *date.Date, categoryID uint) (*entity.Task, error) {
	task, err := s.repo.Create(title, dueDate, categoryID, s.user.ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Edit(id uint, title string, done bool, dueDate *date.Date, categoryID uint) (*entity.Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if task.UserID != s.user.ID {
		return nil, apperror.ErrUnauthorized
	}

	edited, err := s.repo.Edit(id, title, done, dueDate, categoryID)
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

func (s *TaskService) GetTodayTasks() ([]entity.Task, error) {
	tasks, err := s.Get()
	if err != nil {
		return nil, err
	}

	today := date.Now()
	var todayTasks []entity.Task
	for _, task := range tasks {
		if task.DueDate != nil && task.DueDate.IsSameDate(today) {
			todayTasks = append(todayTasks, task)
		}
	}

	return todayTasks, nil
}

func (s *TaskService) GetByDate(date *date.Date) ([]entity.Task, error) {
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

	// edit the task with toggled isDone
	_, err = s.Edit(id, task.Title, isDone, task.DueDate, task.CategoryID)

	return err
}
