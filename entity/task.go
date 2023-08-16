package entity

import "time"

type Task struct {
	Id         uint
	Title      string
	DueDate    time.Time
	CategoryId uint
	Done       bool
}
