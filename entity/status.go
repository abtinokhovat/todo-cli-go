package entity

import "bytes"

type Status struct {
	Category Category
	Tasks    []Task
}

func NewStatus(category Category, task []Task) *Status {
	return &Status{
		Category: category,
		Tasks:    task,
	}
}

func (s Status) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.Category.String())

	if len(s.Tasks) == 0 {
		buffer.WriteString("\n")
		buffer.WriteString("empty!")
	}

	for _, task := range s.Tasks {
		buffer.WriteString("\n")
		buffer.WriteString(task.String())
	}
	buffer.WriteString("\n")

	return buffer.String()
}
