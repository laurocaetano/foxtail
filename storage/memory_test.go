package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryPut(t *testing.T) {
	t.Parallel()

	inMemory := NewInMemory()

	key := "my-key"
	value := []byte("some random value for testing")

	err := inMemory.Put(key, value)

	// Should be able to insert, without errors
	assert.Nil(t, err)

	// Should not allow to insert empty keys
	err = inMemory.Put("", value)
	assert.EqualError(t, err, "The given key cannot be empty")

	// Should not allow to insert nil value keys
	err = inMemory.Put(key, nil)
	assert.EqualError(t, err, "The given value cannot be nil")

	// Should not allow an empty byte array
	err = inMemory.Put(key, []byte{})
	assert.EqualError(t, err, "The given value cannot be empty")
}
