package entity

import (
	"bytes"
	"fmt"
	"todo-cli-go/pkg/date"
)

type Task struct {
	ID         uint       `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	DueDate    *date.Date `json:"due_date,omitempty"`
	Done       bool       `json:"done,omitempty"`
	CategoryID uint       `json:"category_id,omitempty"`
	UserID     uint       `json:"user_id,omitempty"`
}

func NewTask(id uint, title string, done bool, date *date.Date, categoryId, userID uint) *Task {
	return &Task{
		ID:         id,
		Title:      title,
		DueDate:    date,
		Done:       done,
		CategoryID: categoryId,
		UserID:     userID,
	}
}

func (t Task) String() string {
	var buffer bytes.Buffer

	if t.Done {
		buffer.WriteString("[X] - ")
	} else {
		buffer.WriteString("[] - ")
	}

	buffer.WriteString(fmt.Sprintf("#%d-%s", t.ID, t.Title))
	if t.DueDate != nil {
		buffer.WriteString(fmt.Sprintf(" -> %v", t.DueDate))
	}
	return buffer.String()
}
