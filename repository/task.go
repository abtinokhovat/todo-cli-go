package repository

import (
	"sync"
	"todo-cli-go/pkg/date"

	"todo-cli-go/entity"
	"todo-cli-go/error"

	fileHandler "github.com/abtinokhovat/file-handler-go"
)

const taskStoragePath = "storage/todo.json"

var (
	onceTodo               sync.Once
	taskRepositoryInstance *TaskFileRepository
)

type TaskStorageAdapter interface {
	Create(title string, date *date.Date, categoryId, userID uint) (*entity.Task, error)
	Edit(id uint, title string, done bool, date *date.Date, categoryID uint) (*entity.Task, error)
	GetByID(id uint) (*entity.Task, error)
	GetAll() ([]entity.Task, error)
}

type TaskFileRepository struct {
	handler fileHandler.FileIOHandler[entity.Task]
}

func GetTodoFileRepository() *TaskFileRepository {
	onceTodo.Do(func() {
		serializer := fileHandler.NewJsonSerializer[entity.Task]()
		handler := fileHandler.NewJsonIOHandler[entity.Task](taskStoragePath, serializer)
		taskRepositoryInstance = NewTaskFileRepository(handler)
	})

	return taskRepositoryInstance
}
func NewTaskFileRepository(handler fileHandler.FileIOHandler[entity.Task]) *TaskFileRepository {
	return &TaskFileRepository{handler: handler}
}

func (r *TaskFileRepository) Create(title string, date *date.Date, categoryId, userID uint) (*entity.Task, error) {
	id, err := r.newID()
	if err != nil {
		return nil, err
	}

	task := entity.NewTask(id, title, false, date, categoryId, userID)

	err = r.handler.WriteOne(*task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskFileRepository) Edit(id uint, title string, done bool, date *date.Date, categoryID uint) (*entity.Task, error) {
	tasks, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = title
			tasks[i].Done = done
			tasks[i].DueDate = date
			tasks[i].CategoryID = categoryID

			err := r.handler.DeleteAndWrite(tasks)
			if err != nil {
				return nil, err
			}

			return &tasks[i], nil
		}
	}

	return nil, apperror.ErrTaskNotFoundToEdit
}

func (r *TaskFileRepository) GetByID(id uint) (*entity.Task, error) {
	tasks, err := r.GetAll()
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

func (r *TaskFileRepository) GetAll() ([]entity.Task, error) {
	tasks, err := r.handler.Read()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskFileRepository) newID() (uint, error) {
	tasks, err := r.handler.Read()
	if err != nil {
		return 0, err
	}
	return uint(len(tasks) + 1), nil
}
