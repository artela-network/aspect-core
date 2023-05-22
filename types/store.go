package types

type AspectStore interface {

	// Implements KVStore
	Get(prefix, key []byte) []byte
	// Implements KVStore
	Has(prefix, key []byte) bool
	// Implements KVStore
	Set(prefix, key, value []byte)
	// Implements KVStore
	Delete(prfix, key []byte)
}
