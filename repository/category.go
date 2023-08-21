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
	categoryRepositoryInstance *CategoryRepository
)

type CategoryStorageAdapter interface {
	Create(title, color string, userID uint) (*entity.Category, error)
	Edit(id uint, title, color string) (*entity.Category, error)
	GetByID(id uint) (*entity.Category, error)
	GetAll() ([]*entity.Category, error)
}

type CategoryRepository struct {
	handler fileHandler.FileIOHandler[entity.Category]
}

func GetCategoryRepository() *CategoryRepository {
	onceCategory.Do(func() {
		serializer := fileHandler.NewJsonSerializer[entity.Category]()
		handler := fileHandler.NewJsonIOHandler[entity.Category](categoryStoragePath, serializer)
		categoryRepositoryInstance = NewCategoryRepository(handler)
	})

	return categoryRepositoryInstance
}
func NewCategoryRepository(handler fileHandler.FileIOHandler[entity.Category]) *CategoryRepository {
	return &CategoryRepository{handler: handler}
}

func (r *CategoryRepository) Create(title, color string, userId uint) (*entity.Category, error) {
	// generate new id
	id := r.newID()

	// make a new category entity
	category := entity.NewCategory(id, title, color, userId)

	// write category to the storage
	err := r.handler.WriteOne(*category)
	if err != nil {
		return nil, err
	}
	return category, nil
}
func (r *CategoryRepository) GetAll() ([]entity.Category, error) {
	categories, err := r.handler.Read()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
func (r *CategoryRepository) GetByID(id uint) (*entity.Category, error) {
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
func (r *CategoryRepository) Edit(id uint, title, color string) (*entity.Category, error) {
	categories, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		// update category data
		if category.ID == id {
			category.Title = title
			category.Color = color
		}

		// delete all data in file and rewrite it
		err := r.handler.DeleteAndWrite(categories)
		if err != nil {
			return nil, err
		}

		return &category, nil
	}

	return nil, apperror.ErrCategoryNotFoundToEdit
}
func (r *CategoryRepository) newID() uint {
	categories, err := r.handler.Read()
	if err != nil {
		panic("could not get a new id for telegram error happened in reading file")
	}
	return uint(len(categories) + 1)
}
