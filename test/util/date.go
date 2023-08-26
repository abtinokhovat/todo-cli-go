package util_test

import "todo-cli-go/pkg/date"

func GetDate(year, month, day uint) *date.Date {
	d, _ := date.NewDate(year, month, day)
	return d
}
