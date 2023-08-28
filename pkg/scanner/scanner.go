package scanner

import (
	"fmt"
	"strconv"

	"todo-cli-go/error"
)

const NoID = 0

// BufScanner interface to mock bufio.Scanner methods.
type BufScanner interface {
	Scan() bool
	Text() string
	Err() error
}

func Scan(scanner BufScanner, message string) string {
	fmt.Println(message)
	scanner.Scan()
	fmt.Println()
	return scanner.Text()
}

type Scanner struct {
	scanner BufScanner
}

func NewScanner(scanner BufScanner) *Scanner {
	return &Scanner{scanner: scanner}
}

func (s *Scanner) Scan(message string) string {
	return Scan(s.scanner, message)
}
func (s *Scanner) ScanID(message string) (uint, error) {
	scanned := s.Scan(message)
	if scanned == "" {
		return NoID, nil
	}

	id, err := strconv.Atoi(scanned)
	if err != nil {
		return 0, apperror.ErrNotCorrectDigit
	}

	return uint(id), nil
}
