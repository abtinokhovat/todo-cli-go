package repository_test

import "errors"

const (
	ErrOnReading          = "error on reading file"
	ErrOnWriting          = "error on writing file"
	ErrDeleting           = "error on deleting file"
	ErrOnWritingOrReading = "error on deleting and writing file file"
)

type MockIOHandler[T any] struct {
	storage *[]T
	config  MockIOConfig
}

type MockIOConfig struct {
	read   bool
	write  bool
	delete bool
}

func NewMockIOHandler[T any](storage *[]T, config MockIOConfig) *MockIOHandler[T] {
	return &MockIOHandler[T]{
		storage: storage,
		config:  config,
	}
}

func (h *MockIOHandler[T]) Read() ([]T, error) {
	if h.config.read {
		return *h.storage, nil
	}
	return nil, errors.New(ErrOnReading)
}
func (h *MockIOHandler[T]) WriteOne(data T) error {
	if h.config.write {
		*h.storage = append(*h.storage, data)
		return nil
	}
	return errors.New(ErrOnWriting)
}
func (h *MockIOHandler[T]) DeleteAndWrite(data []T) error {
	if h.config.delete && h.config.write {
		h.storage = &data
		return nil
	}
	return errors.New(ErrDeleting)
}
func (h *MockIOHandler[T]) DeleteAll() error {
	if h.config.delete {
		h.storage = new([]T)
		return nil
	}
	return errors.New(ErrOnWritingOrReading)
}
