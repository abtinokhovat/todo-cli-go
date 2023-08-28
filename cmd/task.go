package cmd

import (
	"fmt"
	"todo-cli-go/pkg/date"
	scanner2 "todo-cli-go/pkg/scanner"
	"todo-cli-go/service"
)

type TaskPuppet struct {
	service *service.TaskService
	scanner *scanner2.Scanner
}

func NewTaskPuppet(service *service.TaskService, scanner *scanner2.Scanner) *TaskPuppet {
	return &TaskPuppet{
		service: service,
		scanner: scanner,
	}
}

func (p *TaskPuppet) create() {
	// scan title
	title := p.scanner.Scan("enter your task title")

	// scan due date
	dueDateStr := p.scanner.Scan("enter a due date for your task and if you don't want to add a date,press enter")
	// make a date object from date string
	var dueDate *date.Date = nil

	if dueDateStr != "" {
		d, err := date.NewDateFromString(dueDateStr)
		dueDate = d
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	categoryID, err := p.scanner.ScanID("if you want to assign task to a category enter the category's id or press enter")
	if err != nil {
		fmt.Println(err)
		return
	}

	created, err := p.service.Create(title, dueDate, categoryID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(created.String())

}
func (p *TaskPuppet) list() {
	tasks, err := p.service.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("no task, you can create one")
	}

	for _, task := range tasks {
		fmt.Println(task.String())
	}
}
func (p *TaskPuppet) listToday() {
	tasks, err := p.service.GetTodayTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("no tasks for today, phew")
	}

	for _, task := range tasks {
		fmt.Println(task.String())
	}
}
func (p *TaskPuppet) listByDate() {
	// scan date
	dateStr := p.scanner.Scan("enter a date for searching in tasks")
	// make a date object from date string
	sDate, err := date.NewDateFromString(dateStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks, err := p.service.GetByDate(*sDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("no tasks in the date you requested")
	}

	for _, task := range tasks {
		fmt.Println(task.String())
	}
}
func (p *TaskPuppet) edit() {
	taskId, err := p.scanner.ScanID("enter the id task you want to be edited")
	if err != nil {
		fmt.Println(err)
		return
	}

	title := p.scanner.Scan("enter a title for updating your task and if you dont want to update press enter")
	dueDateStr := p.scanner.Scan("enter a due date for updating your task and if you dont want to update press enter")
	dueDate, err := date.NewDateFromString(dueDateStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	categoryID, err := p.scanner.ScanID("enter a category id for updating your task and if you dont want to update press enter")
	if err != nil {
		fmt.Println(err)
		return
	}

	task := service.TaskUpdate{
		Title:      &title,
		DueDate:    dueDate,
		CategoryID: &categoryID,
	}

	// check if they are empty pass nil in order to skip updating nil fields
	if title == "" {
		task.Title = nil
	}
	if categoryID == 0 {
		task.CategoryID = nil
	}

	edited, err := p.service.Edit(taskId, task)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(edited)

}
func (p *TaskPuppet) toggle() {
	taskID, err := p.scanner.ScanID("enter a task id for toggling")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = p.service.Toggle(taskID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("done.")
}
