// This package provides a high level API for
// inserting, retreving and removing data. The
// implementation will allow different type of
// backend for the storage, such as: `in-memory`
// `disk`, `remote` and etc.
package storage

// High level interface, that enables backend
// implementation for data storage.
type Storage interface {

	// Add a new key/value pair to the storage
	// `key` is a string, that has no limit at the moment.
	//
	// `value` is a `byte` array that is not limited
	// at the moment.
	Put(key string, value []byte) error

	// Gets the value assigned to a given `key`. In
	// case the `key` does not exist, it will return
	// an error.
	Get(key string) ([]byte, error)

	// Deletes and existing key and its associated
	// value. In case the given `key` does not exist,
	// an error is returned instead.
	Delete(key string) error
}
