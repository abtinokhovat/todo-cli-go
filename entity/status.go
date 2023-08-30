package entity

import "bytes"

type Status struct {
	category Category
	tasks    []Task
}

func NewStatus(category Category, task []Task) *Status {
	return &Status{
		category: category,
		tasks:    task,
	}
}

func (s Status) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.category.String())

	if len(s.tasks) == 0 {
		buffer.WriteString("\n")
		buffer.WriteString("empty!")
	}

	for _, task := range s.tasks {
		buffer.WriteString("\n")
		buffer.WriteString(task.String())
	}
	buffer.WriteString("\n")

	return buffer.String()
}
