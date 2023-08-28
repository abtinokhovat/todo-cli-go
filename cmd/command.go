package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo-cli-go/pkg/date"
	"todo-cli-go/repository"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/service"
)

type Command struct {
	scanner *bufio.Scanner
	user    *entity.User

	categoryService *service.CategoryService
	taskService     *service.TaskService
}

func NewCommand(user *entity.User, categoryService *service.CategoryService, taskService *service.TaskService) *Command {
	scanner := bufio.NewScanner(os.Stdin)

	return &Command{
		scanner: scanner,
		user:    user,

		categoryService: categoryService,
		taskService:     taskService,
	}
}

// Define constants for commands
const (
	CreateCategory = "create-category"
	EditCategory   = "edit-category"
	ListCategory   = "list-category"
	CreateTask     = "create-task"
	ListTask       = "list-task"
	ListTaskToday  = "list-task-today"
	ListTaskByDate = "list-task-bydate"
	EditTask       = "edit-task"
	ToggleTask     = "toggle-task"
	Exit           = "exit"
)

func (c *Command) Execute(cmd string) {
	// Check the command and call the appropriate handler
	switch cmd {
	case CreateCategory:
		c.createCategory()
	case EditCategory:
		c.editCategory()
	case ListCategory:
		c.listCategory()
	case CreateTask:
		c.createTask()
	case ListTask:
		c.listTask()
	case ListTaskToday:
		c.listTaskToday()
	case ListTaskByDate:
		c.listTaskByDate()
	case EditTask:
		c.editTask()
	case ToggleTask:
		c.toggleTask()
	case Exit:
		os.Exit(1)
	default:
		fmt.Println("Invalid command")
	}
}

// Define command-specific handler methods
func (c *Command) createCategory() {

	title := Scan(c.scanner, "enter a title for your new category")
	color := Scan(c.scanner, "enter a color for your new category")

	category, err := c.categoryService.Create(title, color)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(category.String())
}
func (c *Command) editCategory() {
	id, err := strconv.Atoi(Scan(c.scanner, "enter the id of category you want to edit"))
	if err != nil {
		fmt.Println(apperror.ErrNotCorrectDigit)
		return
	}

	title := Scan(c.scanner, "enter a title for updating")
	color := Scan(c.scanner, "enter a color for updating")

	category, err := c.categoryService.Edit(uint(id), title, color)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("update successful")
	fmt.Println(category)
}
func (c *Command) listCategory() {
	categories, err := c.categoryService.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(categories) == 0 {
		fmt.Println("no categories :( ,make one")
	}

	for _, category := range categories {
		fmt.Println(category.String())
	}
}

func (c *Command) createTask() {
	// scan title
	title := c.scan("enter your task title")

	// scan due date
	dueDateStr := c.scan("enter a due date for your task and if you don't want to add a date,press enter")
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

	categoryID, err := c.scanID("if you want to assign task to a category enter the category's id or press enter")
	if err != nil {
		fmt.Println(err)
		return
	}

	created, err := c.taskService.Create(title, dueDate, categoryID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(created.String())

}
func (c *Command) listTask() {
	tasks, err := c.taskService.Get()
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
func (c *Command) listTaskToday() {
	tasks, err := c.taskService.GetTodayTasks()
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
func (c *Command) listTaskByDate() {
	// scan date
	dateStr := c.scan("enter a date for searching in tasks")
	// make a date object from date string
	sDate, err := date.NewDateFromString(dateStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks, err := c.taskService.GetByDate(*sDate)
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
func (c *Command) editTask() {
	taskId, err := c.scanID("enter the id task you want to be edited")
	if err != nil {
		fmt.Println(err)
		return
	}

	title := c.scan("enter a title for updating your task and if you dont want to update press enter")
	dueDateStr := c.scan("enter a due date for updating your task and if you dont want to update press enter")
	dueDate, err := date.NewDateFromString(dueDateStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	categoryID, err := c.scanID("enter a category id for updating your task and if you dont want to update press enter")
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

	edited, err := c.taskService.Edit(taskId, task)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(edited)

}
func (c *Command) toggleTask() {
	taskID, err := c.scanID("enter a task id for toggling")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.taskService.Toggle(taskID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("done.")
}

func (c *Command) scan(message string) string {
	return Scan(c.scanner, message)
}
func (c *Command) scanID(message string) (uint, error) {
	scanned := c.scan(message)
	if scanned == "" {
		return repository.NoCategory, nil
	}

	id, err := strconv.Atoi(scanned)
	if err != nil {
		return 0, apperror.ErrNotCorrectDigit
	}

	return uint(id), nil
}
