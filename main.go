package main

import (
	"bufio"
	"fmt"
	"os"

	"todo-cli-go/cmd"
	"todo-cli-go/entity"
	"todo-cli-go/service"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	authService := service.BuildUserService()

	entityType, operation := parseArgs(os.Args)

	user := handleAuthentication(scanner, authService, entityType, operation)
	if user == nil {
		return
	}

	initCommand := cmd.CommandBuilder(operation, entityType)

	categoryService := service.BuildCategoryService(user)

	command := cmd.NewCommand(user, categoryService)
	runCommandLoop(command, initCommand, scanner)
}

func parseArgs(args []string) (string, string) {
	entityType := "task"
	operation := "list"

	if len(args) > 1 {
		entityType = args[1]
	}
	if len(args) > 2 {
		operation = args[2]
	}

	return entityType, operation
}

func handleAuthentication(scanner *bufio.Scanner, authService *service.UserService, entityType, operation string) *entity.User {
	if operation == "register" && entityType == "user" {
		return registerAndLogin(scanner, authService)
	}

	return login(scanner, authService)
}

func login(scanner *bufio.Scanner, authService *service.UserService) *entity.User {
	email := cmd.Scan(scanner, "Enter your email address")
	password := cmd.Scan(scanner, "Enter your password")

	user, err := authService.Login(email, password)
	if err != nil {
		fmt.Println(err)
	}

	return user
}

func registerAndLogin(scanner *bufio.Scanner, authService *service.UserService) *entity.User {
	email := cmd.Scan(scanner, "Enter your email address")
	password := cmd.Scan(scanner, "Enter your password")

	_, err := authService.Register(email, password)
	if err != nil {
		fmt.Println("Error on registering user:", err)
	}

	user, err := authService.Login(email, password)
	if err != nil {
		fmt.Println("Error on logging in user:", err)
	}

	return user
}

func runCommandLoop(command *cmd.Command, initCommand string, scanner *bufio.Scanner) {
	if initCommand != "register-user" {
		command.Execute(initCommand)
	}

	for {
		commandString := cmd.Scan(scanner, "Enter the command you want to execute")
		command.Execute(commandString)
	}
}
