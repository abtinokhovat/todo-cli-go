package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"todo-cli-go/entity"
	"todo-cli-go/error"
	"todo-cli-go/service"
)

type Command struct {
	scanner         *bufio.Scanner
	user            *entity.User
	categoryService *service.CategoryService
}

func NewCommand(user *entity.User, categoryService *service.CategoryService) *Command {
	scanner := bufio.NewScanner(os.Stdin)

	return &Command{
		user:            user,
		categoryService: categoryService,
		scanner:         scanner,
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
	fmt.Println("create-task")
}
func (c *Command) listTask() {
	fmt.Println("list-task")
}
func (c *Command) listTaskToday() {
	fmt.Println("list-task-today")
}
func (c *Command) listTaskByDate() {
	fmt.Println("list-task-bydate")
}
func (c *Command) editTask() {
	fmt.Println("edit-task")
}
func (c *Command) toggleTask() {
	fmt.Println("toggle-task")
}
