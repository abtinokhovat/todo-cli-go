package entity

import (
	"todo-cli-go/color"
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
	return color.Colorf(c.colorCode, "#%d-%s: 🚻%d", c.ID, c.Title, c.UserID)
}
