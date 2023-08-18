package entity

import "time"

type Task struct {
	Id         uint
	Title      string
	DueDate    time.Time
	Done       bool
	CategoryId uint
	UserId     uint
}
