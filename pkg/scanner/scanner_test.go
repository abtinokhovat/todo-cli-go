package scanner_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-cli-go/error"
	"todo-cli-go/pkg/scanner"
)

func TestScanner_ScanID(t *testing.T) {
	testCases := []struct {
		name        string
		mockInput   string
		expectedID  uint
		expectedErr error
	}{
		{
			name:        "ordinary scan valid ID",
			mockInput:   "123",
			expectedID:  123,
			expectedErr: nil,
		},
		{
			name:        "scan with empty input",
			mockInput:   "",
			expectedID:  scanner.NoID,
			expectedErr: nil,
		},
		{
			name:        "scan with non-numeric input",
			mockInput:   "abc",
			expectedID:  0,
			expectedErr: apperror.ErrNotCorrectDigit,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			mock := &mockScanner{data: tc.mockInput}
			s := scanner.NewScanner(mock)

			// Execution
			id, err := s.ScanID("Enter an ID: ")

			// Assertion
			assert.Equal(t, tc.expectedID, id)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.expectedErr))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// mockScanner is a mock implementation of Scanner.
type mockScanner struct {
	data string
}

// Scan simulates the Scan method for the mock scanner.
func (m *mockScanner) Scan() bool {
	return true
}

// Text simulates the Text method for the mock scanner.
func (m *mockScanner) Text() string {
	return m.data
}

// Err simulates the Err method for the mock scanner.
func (m *mockScanner) Err() error {
	return nil
}
