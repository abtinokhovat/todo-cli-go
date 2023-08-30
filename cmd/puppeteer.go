package cmd

import (
	"bufio"
	"fmt"
	"os"

	scanner2 "todo-cli-go/pkg/scanner"
	"todo-cli-go/service"
)

type Puppeteer struct {
	categoryPuppet *CategoryPuppet
	taskPuppet     *TaskPuppet
	statusPuppet   *StatusPuppet
}

func NewPuppeteer(statusMaster service.StatusMaster, categoryMaster service.CategoryMaster, taskMaster service.TaskMaster) *Puppeteer {
	scanner := scanner2.NewScanner(bufio.NewScanner(os.Stdin))

	return &Puppeteer{
		statusPuppet:   NewStatusPuppet(statusMaster),
		categoryPuppet: NewCategoryPuppet(categoryMaster, scanner),
		taskPuppet:     NewTaskPuppet(taskMaster, scanner),
	}
}

// Define constants for commands
const (
	AllStatus      = "status-all"
	DoneStatus     = "done-status"
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

func (c *Puppeteer) Execute(cmd string) {
	// Check the command and call the appropriate handler
	switch cmd {
	case AllStatus:
		c.statusPuppet.Status()
	case DoneStatus:
		c.statusPuppet.Done()
	case CreateCategory:
		c.categoryPuppet.create()
	case EditCategory:
		c.categoryPuppet.edit()
	case ListCategory:
		c.categoryPuppet.list()
	case CreateTask:
		c.taskPuppet.create()
	case ListTask:
		c.taskPuppet.list()
	case ListTaskToday:
		c.taskPuppet.listToday()
	case ListTaskByDate:
		c.taskPuppet.listByDate()
	case EditTask:
		c.taskPuppet.edit()
	case ToggleTask:
		c.taskPuppet.toggle()
	case Exit:
		os.Exit(1)
	default:
		fmt.Println("Invalid command")
	}
}
