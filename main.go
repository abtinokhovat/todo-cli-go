package main

import (
	"bufio"
	"os"

	"todo-cli-go/cmd"
	"todo-cli-go/entity"
	"todo-cli-go/pkg/scanner"
	"todo-cli-go/service"
)

func main() {
	// build a scanner for entire application
	scn := scanner.NewScanner(bufio.NewScanner(os.Stdin))

	// parse entity and operation
	entityType, operation := parseArgs(os.Args)

	// build the auth service
	authService := service.BuildUserService()
	// build the auth puppet
	authPuppet := cmd.NewAuthPuppet(authService, scn)

	// handle login and register flow and don't let the user go further if login failed
	user := handleAuthentication(authPuppet, entityType, operation)
	if user == nil {
		return
	}

	// run the init command
	initCommand := cmd.CommandBuilder(operation, entityType)

	// build task and category services
	categoryService := service.BuildCategoryService(user)
	categoryValidationService := service.NewCategoryValidationService(categoryService)

	taskService := service.BuildTaskService(user, categoryValidationService)

	statusService := service.NewStatusService(categoryService, taskService)

	// build the master puppet for directing the puppets
	command := cmd.NewPuppeteer(statusService, categoryService, taskService)

	// running application loop
	if initCommand != "register-user" {
		command.Execute(initCommand)
	}

	for {
		commandString := scn.Scan("Enter the command you want to execute")
		command.Execute(commandString)
	}
}

func parseArgs(args []string) (string, string) {
	entityType := "status"
	operation := "done"

	if len(args) > 1 {
		entityType = args[1]
	}
	if len(args) > 2 {
		operation = args[2]
	}

	return entityType, operation
}

func handleAuthentication(authPuppet *cmd.AuthPuppet, entityType, operation string) *entity.User {
	if operation == "register" && entityType == "user" {
		return authPuppet.RegisterAndLogin()
	}

	return authPuppet.Login()
}
