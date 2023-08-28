package scanner

import (
	"bufio"
	"strconv"
	"todo-cli-go/cmd"

	"todo-cli-go/error"
	"todo-cli-go/repository"
)

type Scanner struct {
	scanner *bufio.Scanner
}

func NewScanner(scanner *bufio.Scanner) *Scanner {
	return &Scanner{scanner: scanner}
}

func (s *Scanner) Scan(message string) string {
	return cmd.Scan(s.scanner, message)
}
func (s *Scanner) ScanID(message string) (uint, error) {
	scanned := s.Scan(message)
	if scanned == "" {
		return repository.NoCategory, nil
	}

	id, err := strconv.Atoi(scanned)
	if err != nil {
		return 0, apperror.ErrNotCorrectDigit
	}

	return uint(id), nil
}
