package repository

import (
	fileHandler "github.com/abtinokhovat/file-handler-go"
	"sync"
	"todo-cli-go/entity"
)

const categoryStoragePath = "storage/category.json"

var (
	onceCategory               sync.Once
	categoryRepositoryInstance *CategoryRepository
)

type CategoryStorageAdapter interface {
	Create(userID int, title, color string) (*entity.Category, error)
	Edit(id int, title, color string) (*entity.Category, error)
	GetByID(id int) (*entity.Category, error)
	GetAll() []*entity.Category
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

func (r *CategoryRepository) Create(userId uint, title, color string) (*entity.Category, error) {
	id := r.newID()
	category := entity.NewCategory(id, title, color, userId)
	r.handler.WriteOne()
	//TODO implement me
	panic("implement me")
}
func (r *CategoryRepository) Edit(id int, title, color string) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}
func (r *CategoryRepository) GetByID(id int) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}
func (r *CategoryRepository) GetAll() []*entity.Category {
	//TODO implement me
	panic("implement me")
}
func (r *CategoryRepository) newID() uint {
	//TODO implement me
	panic("implement me")
}
