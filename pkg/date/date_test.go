package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	apperror "todo-cli-go/error"
)

func TestNewDate(t *testing.T) {
	testCases := []struct {
		name string
		date Date
		err  error
	}{
		{
			name: "ordinary make should make a Date",
			date: Date{year: 2023, month: 8, day: 28},
		},
		{
			name: "values over validation limit should have error",
			date: Date{year: 2023, month: 123, day: 44},
			err:  apperror.ErrNotInACorrectFormat,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 2. execution
			date, err := NewDate(tc.date.year, tc.date.month, tc.date.day)

			// 3. assertion
			if tc.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.date.year, date.year)
				assert.Equal(t, tc.date.month, date.month)
				assert.Equal(t, tc.date.day, date.day)
			}
		})
	}
}

func TestNewDateFromString(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		expected *Date
		err      error
	}{
		{
			name:     "ordinary make with string should make a Date",
			str:      "2023-08-28",
			expected: &Date{year: 2023, month: 8, day: 28},
		},
		{
			name:     "string of \"\" should make an empty Date",
			str:      "",
			expected: nil,
		},
		{
			name:     "string should be in the correct format",
			str:      "2023/08/28",
			expected: nil,
			err:      apperror.ErrNotInACorrectFormat,
		},
		{
			name:     "year and month and day have to be digits",
			str:      "2023-08-28a",
			expected: nil,
			err:      apperror.ErrNotCorrectDigit,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 2. execution
			date, err := NewDateFromString(tc.str)

			// 3. assertion
			if tc.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, date)
			}
		})
	}
}

func TestValidator(t *testing.T) {
	testCases := []struct {
		name   string
		year   uint
		month  uint
		day    uint
		result bool
	}{
		{
			name:   "should validate date values",
			year:   2023,
			month:  8,
			day:    28,
			result: true,
		},
		{
			name:   "year, month, and day should be valid",
			year:   2023,
			month:  13,
			day:    32,
			result: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 2. execution
			valid := Validator(tc.year, tc.month, tc.day)

			// 3. assertion
			assert.Equal(t, tc.result, valid)
		})
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		name     string
		date     Date
		expected string
	}{
		{
			name:     "should make a string with the correct format",
			date:     Date{year: 2023, month: 8, day: 28},
			expected: "2023-08-28",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 2. execution
			str := tc.date.String()

			// 3. assertion
			assert.Equal(t, tc.expected, str)
		})
	}
}

func TestIsSameDate(t *testing.T) {
	testCases := []struct {
		name     string
		date     Date
		other    Date
		expected bool
	}{
		{
			name:     "should return true if both dates were the same",
			date:     Date{year: 2023, month: 8, day: 28},
			other:    Date{year: 2023, month: 8, day: 28},
			expected: true,
		},
		{
			name:     "should return false if they were not on the same date",
			date:     Date{year: 2023, month: 8, day: 28},
			other:    Date{year: 2023, month: 8, day: 29},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 2. execution
			result := tc.date.IsSameDate(tc.other)

			// 3. assertion
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		date     Date
		expected string
	}{
		{
			name:     "should marshal to the correct JSON string",
			date:     Date{year: 2023, month: 8, day: 28},
			expected: "\"2023-08-28\"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 2. execution
			data, err := tc.date.MarshalJSON()

			// 3. assertion
			assert.NoError(t, err)
			assert.Equal(t, []byte(tc.expected), data)
		})
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected Date
	}{
		{
			name:     "should unmarshal from the correct JSON string",
			data:     []byte("\"2023-08-28\""),
			expected: Date{year: 2023, month: 8, day: 28},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 2. execution
			var date Date
			err := date.UnmarshalJSON(tc.data)

			// 3. assertion
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, date)
		})
	}
}
