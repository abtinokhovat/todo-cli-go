package service

import (
	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/repository"
)

type CategoryMaster interface {
	Create(title, color string) (*entity.Category, error)
	Get() ([]entity.Category, error)
	Edit(id uint, title, color string) (*entity.Category, error)
}

type CategoryService struct {
	user *entity.User
	repo repository.CategoryStorageAdapter
}

func BuildCategoryService(user *entity.User) *CategoryService {
	repo := repository.GetCategoryFileRepository()
	return NewCategoryService(user, repo)
}

func NewCategoryService(user *entity.User, repository repository.CategoryStorageAdapter) *CategoryService {
	return &CategoryService{user: user, repo: repository}
}

func (s *CategoryService) Create(title, color string) (*entity.Category, error) {
	category, err := s.repo.Create(title, color, s.user.ID)
	if err != nil {
		return nil, err
	}

	return category, err
}
func (s *CategoryService) Get() ([]entity.Category, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var userCategories []entity.Category
	for _, category := range categories {
		if category.UserID == s.user.ID {
			userCategories = append(userCategories, category)
		}
	}

	return userCategories, nil
}
func (s *CategoryService) Edit(id uint, title, color string) (*entity.Category, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// if the category was not for the user raise error
	if category.UserID != s.user.ID {
		return nil, apperror.ErrUnauthorized
	}

	edited, err := s.repo.Edit(id, title, color)
	if err != nil {
		return nil, err
	}

	return edited, nil
}
