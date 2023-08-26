package repository_test

import (
	apperror "todo-cli-go/error"
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
	return nil, apperror.ErrOnReading
}
func (h *MockIOHandler[T]) WriteOne(data T) error {
	if h.config.write {
		*h.storage = append(*h.storage, data)
		return nil
	}
	return apperror.ErrOnWriting
}
func (h *MockIOHandler[T]) DeleteAndWrite(data []T) error {
	if h.config.delete && h.config.write {
		h.storage = &data
		return nil
	}
	return apperror.ErrOnWritingOrReading
}
func (h *MockIOHandler[T]) DeleteAll() error {
	if h.config.delete {
		h.storage = new([]T)
		return nil
	}
	return apperror.ErrOnWritingOrReading
}
