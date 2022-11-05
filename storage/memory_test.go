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

func TestInMemoryGet(t *testing.T) {
	t.Parallel()

	inMemory := NewInMemory()

	// Empty key not allowed
	val, err := inMemory.Get("")
	assert.Nil(t, val)
	assert.EqualError(t, err, "The given key cannot be empty")

	// Valid key
	key := "fox"
	value := []byte("tail")
	inMemory.Put(key, value)

	ret, err := inMemory.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, value, ret)

	// Non existing key
	ret, err = inMemory.Get("seal corp.")
	assert.Nil(t, err)
	assert.Nil(t, ret)
}
