package cmd

import (
	"errors"
	"fmt"
	"todo-cli-go/entity"
	"todo-cli-go/pkg/scanner"
	"todo-cli-go/service"
)

type AuthPuppet struct {
	scanner *scanner.Scanner
	service *service.UserService
}

func NewAuthPuppet(service *service.UserService, scanner *scanner.Scanner) *AuthPuppet {
	return &AuthPuppet{
		scanner: scanner,
		service: service,
	}
}

func (p *AuthPuppet) Login() *entity.User {
	email := p.scanner.Scan("Enter your email address")
	password := p.scanner.Scan("Enter your password")

	user, err := p.service.Login(email, password)
	if err != nil {
		fmt.Println(err)
	}

	return user
}

func (p *AuthPuppet) Register() (string, string, error) {
	email := p.scanner.Scan("Enter your email address")
	password := p.scanner.Scan("Enter your password")

	_, err := p.service.Register(email, password)
	if err != nil {
		return "", "", errors.New(fmt.Sprintln("Error on registering user:", err))
	}

	return email, password, nil
}

func (p *AuthPuppet) RegisterAndLogin() *entity.User {
	email, password, err := p.Register()
	if err != nil {
		fmt.Println(err)
	}

	user, err := p.service.Login(email, password)
	if err != nil {
		fmt.Println("Error on logging in user:", err)
	}

	return user
}
