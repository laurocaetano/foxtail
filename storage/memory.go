package storage

import (
	"errors"

	"github.com/orcaman/concurrent-map/v2"
)

// Struct that will implement the Storage interface
// and have an in-memory map as its backend storage.
type InMemory struct {
	backend cmap.ConcurrentMap[string, []byte]
}

func NewInMemory() InMemory {
	return InMemory{
		backend: cmap.New[[]byte](),
	}
}

func (memory *InMemory) Put(key string, value []byte) error {
	if key == "" {
		return errors.New("The given key cannot be empty")
	}

	if value == nil {
		return errors.New("The given value cannot be nil")
	}

	if len(value) < 1 {
		return errors.New("The given value cannot be empty")
	}

	memory.backend.Set(key, value)

	return nil
}

func (memory *InMemory) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, errors.New("The given key cannot be empty")
	}

	val, _ := memory.backend.Get(key)

	return val, nil
}

func (memory *InMemory) Delete(key string) error {
	if key == "" {
		return errors.New("The given key cannot be empty")
	}

	memory.backend.Remove(key)

	// So far since deleting from a map is a no-op, no error
	// is ever returned
	return nil
}
