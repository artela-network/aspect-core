package types

type AspectStore interface {
	// Implements KVStore
	Get(key []byte) []byte
	// Implements KVStore
	Has(key []byte) bool
	// Implements KVStore
	Set(key, value []byte)
	// Implements KVStore
	Delete(key []byte)
}
