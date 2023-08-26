package date

import (
	"errors"
	"fmt"
)

type Date struct {
	year  uint
	month uint
	day   uint
}

func NewDate(year, month, day uint) (*Date, error) {
	if !Validator(year, month, day) {
		return nil, errors.New("not correct format")
	}

	return &Date{
		year:  year,
		month: month,
		day:   day,
	}, nil
}

func Validator(year, month, day uint) bool {
	if year < 0 || day > 31 || month > 13 {
		return false
	}
	return true
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%d-%d", d.year, d.day, d.month)
}
