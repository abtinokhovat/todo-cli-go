package entity

import (
	"bytes"
	"fmt"
	"todo-cli-go/pkg/date"
)

type Task struct {
	ID         uint
	Title      string
	DueDate    *date.Date
	Done       bool
	CategoryID uint
	UserId     uint
}

func NewTask(id uint, title string, done bool, date *date.Date, categoryId, userID uint) *Task {
	return &Task{
		ID:         id,
		Title:      title,
		DueDate:    date,
		Done:       done,
		CategoryID: categoryId,
		UserId:     userID,
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
