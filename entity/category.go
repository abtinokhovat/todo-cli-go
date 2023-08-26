package entity

import (
	"todo-cli-go/pkg/color"
)

type Category struct {
	ID        uint
	Title     string
	Color     string
	UserID    uint
	colorCode color.Color
}

func NewCategory(id uint, title, colour string, userId uint) *Category {
	colorCode := color.GetColor(colour)

	return &Category{
		ID:        id,
		Title:     title,
		Color:     colour,
		UserID:    userId,
		colorCode: colorCode,
	}
}

func (c *Category) String() string {
	// add color code if it was not initialized
	if c.colorCode == "" {
		c.colorCode = color.GetColor(c.Color)
	}
	return color.Colorf(c.colorCode, "#%d-%s", c.ID, c.Title)
}
