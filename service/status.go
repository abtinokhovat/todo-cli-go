package service

import (
	"todo-cli-go/entity"
	"todo-cli-go/pkg/scanner"
)

type StatusMaster interface {
	GetOverall() ([]entity.Status, error)
}

type StatusService struct {
	categoryMaster CategoryMaster
	taskMaster     TaskMaster
}

func NewStatusService(category CategoryMaster, task TaskMaster) StatusMaster {
	return &StatusService{
		categoryMaster: category,
		taskMaster:     task,
	}
}

func (s StatusService) GetOverall() ([]entity.Status, error) {
	// get categories of user
	categories, err := s.categoryMaster.Get()
	if err != nil {
		return nil, err
	}

	// get all tasks of the user
	allTasks, err := s.taskMaster.Get()
	if err != nil {
		return nil, err
	}

	// group tasks which don't have a category into a status struct
	var tasksWithoutCategory []entity.Task
	notACategory := entity.Category{ID: scanner.NoID}

	for _, task := range allTasks {
		if task.CategoryID == scanner.NoID {
			tasksWithoutCategory = append(tasksWithoutCategory, task)
		}
	}

	noCategoryStatus := entity.NewStatus(notACategory, tasksWithoutCategory)

	// group other tasks which have category

	var stats []entity.Status

	for _, category := range categories {
		var t []entity.Task
		for _, task := range allTasks {
			if task.CategoryID == category.ID {
				t = append(t, task)
			}
		}

		stat := entity.NewStatus(category, t)
		stats = append(stats, *stat)
	}

	stats = append(stats, *noCategoryStatus)

	return stats, nil
}
