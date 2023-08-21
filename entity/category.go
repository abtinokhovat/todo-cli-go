package entity

type Category struct {
	ID     uint
	Title  string
	Color  string
	UserID uint
}

func NewCategory(id uint, title, color string, userId uint) *Category {
	return &Category{
		ID:     id,
		Title:  title,
		Color:  color,
		UserID: userId,
	}
}
