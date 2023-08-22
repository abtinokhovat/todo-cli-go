package repository_test

import "errors"

const (
	ErrOnReading          = "error on reading file"
	ErrWriting            = "error on writing file"
	ErrDeleting           = "error on deleting file"
	ErrOnWritingOrReading = "error on deleting and writing file file"
)

type MockIOHandler[T any] struct {
	storage *[]T
}

func NewMockIOHandler[T any](storage *[]T) *MockIOHandler[T] {
	return &MockIOHandler[T]{
		storage: storage,
	}
}

func (h *MockIOHandler[T]) Read() ([]T, error) {
	return *h.storage, nil
}
func (h *MockIOHandler[T]) WriteOne(data T) error {
	*h.storage = append(*h.storage, data)
	return nil
}
func (h *MockIOHandler[T]) DeleteAndWrite(data []T) error {
	h.storage = &data
	return nil
}
func (h *MockIOHandler[T]) DeleteAll() error {
	h.storage = new([]T)
	return nil
}

type MockBadIOHandler[T any] struct {
}

func (m MockBadIOHandler[T]) Read() ([]T, error) {
	return nil, errors.New(ErrOnReading)
}

func (m MockBadIOHandler[T]) WriteOne(data T) error {
	return errors.New(ErrWriting)
}

func (m MockBadIOHandler[T]) DeleteAll() error {
	return errors.New(ErrDeleting)
}

func (m MockBadIOHandler[T]) DeleteAndWrite(data []T) error {
	return errors.New(ErrOnWritingOrReading)
}
