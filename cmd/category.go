package cmd

import (
	"fmt"

	"todo-cli-go/error"
	"todo-cli-go/pkg/scanner"
	"todo-cli-go/service"
)

type CategoryPuppet struct {
	service service.CategoryMaster
	scanner *scanner.Scanner
}

func NewCategoryPuppet(service service.CategoryMaster, scanner *scanner.Scanner) *CategoryPuppet {
	return &CategoryPuppet{
		service: service,
		scanner: scanner,
	}
}

func (p *CategoryPuppet) create() {

	title := p.scanner.Scan("enter a title for your new category")
	color := p.scanner.Scan("enter a color for your new category")

	category, err := p.service.Create(title, color)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(category.String())
}
func (p *CategoryPuppet) edit() {
	id, err := p.scanner.ScanID("enter the id of category you want to edit")
	if err != nil {
		fmt.Println(apperror.ErrNotCorrectDigit)
		return
	}

	title := p.scanner.Scan("enter a title for updating")
	color := p.scanner.Scan("enter a color for updating")

	category, err := p.service.Edit(id, title, color)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("update successful")
	fmt.Println(category)
}
func (p *CategoryPuppet) list() {
	categories, err := p.service.Get()
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
