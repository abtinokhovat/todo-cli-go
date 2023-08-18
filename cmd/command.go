package cmd

import (
	"fmt"
	"os"
	"todo-cli-go/entity"
)

type Command struct {
	user *entity.User
}

func NewCommand(user *entity.User) *Command {
	return &Command{
		user: user,
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
	fmt.Println("create-category")
}
func (c *Command) editCategory() {
	fmt.Println("edit-category")
}
func (c *Command) listCategory() {
	fmt.Println("list-category")
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
