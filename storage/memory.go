package storage

import "errors"

// Struct that will implement the Storage interface
// and have an in-memory map as its backend storage.
type InMemory struct {
	backend map[string][]byte
}

func NewInMemory() InMemory {
	return InMemory{
		backend: make(map[string][]byte),
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

	// To start with, a non thread-safe approach
	memory.backend[key] = value

	return nil
}

func (memory *InMemory) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, errors.New("The given key cannot be empty")
	}

	val, exists := memory.backend[key]

	if !exists {
		return nil, nil
	}

	return val, nil
}
