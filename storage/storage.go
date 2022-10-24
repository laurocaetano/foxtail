// This package provides a high level API for
// inserting, retreving and removing data. The
// implementation will allow different type of
// backend for the storage, such as: `in-memory`
// `disk`, `remote` and etc.
package storage

type Storage interface {
	Put(key string, value byte) error
	Get(key string) (byte, error)
	Delete(key string) error
}
