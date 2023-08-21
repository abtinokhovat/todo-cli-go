package repository_test

type MockIOHandler[T any] struct {
	storage *[]T
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
