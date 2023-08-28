package date

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"todo-cli-go/error"
)

type Date struct {
	year  uint
	month uint
	day   uint
}

func NewDate(year, month, day uint) (*Date, error) {
	if !Validator(year, month, day) {
		return nil, apperror.ErrNotInACorrectFormat
	}

	return &Date{
		year:  year,
		month: month,
		day:   day,
	}, nil
}
func NewDateFromString(str string) (*Date, error) {
	if str == "" {
		return nil, nil
	}

	data := strings.Split(str, "-")

	if len(data) != 3 {
		return nil, apperror.ErrNotInACorrectFormat
	}

	parsed := make([]uint, 3)
	for i, value := range data {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, apperror.ErrNotCorrectDigit
		}
		parsed[i] = uint(num)
	}

	year := parsed[0]
	month := parsed[1]
	day := parsed[2]

	return NewDate(year, month, day)
}

func Validator(year, month, day uint) bool {
	if year < 0 || day > 31 || month > 13 {
		return false
	}
	return true
}

func Now() *Date {
	now := time.Now()
	year := uint(now.Year())
	month := uint(now.Month())
	day := uint(now.Day())

	return &Date{
		year:  year,
		month: month,
		day:   day,
	}
}

func (d *Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.year, d.month, d.day)
}
func (d *Date) IsSameDate(other Date) bool {
	return d.year == other.year && d.month == other.month && d.day == other.day
}
func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}
func (d *Date) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.ReplaceAll(str, "\"", "")
	parts := strings.Split(str, "-")

	if len(parts) != 3 {
		return errors.New("date not in correct format")
	}

	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	d.year = uint(year)
	d.month = uint(month)
	d.day = uint(day)

	return nil
}
