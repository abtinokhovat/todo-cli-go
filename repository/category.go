package repository

import (
	fileHandler "github.com/abtinokhovat/file-handler-go"
	"sync"
	"todo-cli-go/entity"
	apperror "todo-cli-go/error"
)

const categoryStoragePath = "storage/category.json"

var (
	onceCategory               sync.Once
	categoryRepositoryInstance *CategoryFileRepository
)

type CategoryStorageAdapter interface {
	Create(title, color string, userID uint) (*entity.Category, error)
	Edit(id uint, title, color string) (*entity.Category, error)
	GetByID(id uint) (*entity.Category, error)
	GetAll() ([]entity.Category, error)
}

type CategoryFileRepository struct {
	handler fileHandler.FileIOHandler[entity.Category]
}

func GetCategoryFileRepository() *CategoryFileRepository {
	onceCategory.Do(func() {
		serializer := fileHandler.NewJsonSerializer[entity.Category]()
		handler := fileHandler.NewJsonIOHandler[entity.Category](categoryStoragePath, serializer)
		categoryRepositoryInstance = NewCategoryFileRepository(handler)
	})

	return categoryRepositoryInstance
}
func NewCategoryFileRepository(handler fileHandler.FileIOHandler[entity.Category]) *CategoryFileRepository {
	return &CategoryFileRepository{handler: handler}
}

func (r *CategoryFileRepository) Create(title, color string, userId uint) (*entity.Category, error) {
	// generate new id
	id, err := r.newID()
	if err != nil {
		return nil, err
	}
	// make a new category entity
	category := entity.NewCategory(id, title, color, userId)

	// write category to the storage
	err = r.handler.WriteOne(*category)
	if err != nil {
		return nil, err
	}
	return category, nil
}
func (r *CategoryFileRepository) GetAll() ([]entity.Category, error) {
	categories, err := r.handler.Read()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
func (r *CategoryFileRepository) GetByID(id uint) (*entity.Category, error) {
	categories, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		if category.ID == id {
			return &category, nil
		}
	}

	return nil, apperror.ErrCategoryNotFound
}
func (r *CategoryFileRepository) Edit(id uint, title, color string) (*entity.Category, error) {
	categories, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for i, _ := range categories {
		// update category data
		if categories[i].ID == id {
			categories[i].Title = title
			categories[i].Color = color
		}

		// delete all data in file and rewrite it
		err := r.handler.DeleteAndWrite(categories)
		if err != nil {
			return nil, err
		}

		return &categories[i], nil
	}

	return nil, apperror.ErrCategoryNotFoundToEdit
}
func (r *CategoryFileRepository) newID() (uint, error) {
	categories, err := r.handler.Read()
	if err != nil {
		return 0, err
	}
	return uint(len(categories) + 1), nil
}
