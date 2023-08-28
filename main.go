package main

import (
	"bufio"
	"fmt"
	"os"

	"todo-cli-go/cmd"
	"todo-cli-go/entity"
	"todo-cli-go/pkg/scanner"
	"todo-cli-go/service"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)

	entityType, operation := parseArgs(os.Args)

	authService := service.BuildUserService()

	user := handleAuthentication(authService, entityType, operation, scn)
	if user == nil {
		return
	}

	initCommand := cmd.CommandBuilder(operation, entityType)

	categoryService := service.BuildCategoryService(user)
	taskService := service.BuildTaskService(user)

	command := cmd.NewPuppeteer(categoryService, taskService)

	// running application loop
	if initCommand != "register-user" {
		command.Execute(initCommand)
	}

	for {
		commandString := scanner.Scan(scn, "Enter the command you want to execute")
		command.Execute(commandString)
	}
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

func handleAuthentication(authService service.AuthService, entityType, operation string, scanner *bufio.Scanner) *entity.User {
	if operation == "register" && entityType == "user" {
		return registerAndLogin(authService, scanner)
	}

	return login(authService, scanner)
}

func login(authService service.AuthService, scn *bufio.Scanner) *entity.User {
	email := scanner.Scan(scn, "Enter your email address")
	password := scanner.Scan(scn, "Enter your password")

	user, err := authService.Login(email, password)
	if err != nil {
		fmt.Println(err)
	}

	return user
}

func registerAndLogin(authService service.AuthService, scn *bufio.Scanner) *entity.User {
	email := scanner.Scan(scn, "Enter your email address")
	password := scanner.Scan(scn, "Enter your password")

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
