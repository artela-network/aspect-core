package types

type AspectStore interface {
	//// Implements KVStore
	//Get(prefix, key []byte) []byte
	//// Implements KVStore
	//Has(prefix, key []byte) bool
	//// Implements KVStore
	//Set(prefix, key, value []byte)
	//// Implements KVStore
	//Delete(prfix, key []byte)

	Get(key []byte) []byte

	// Has checks if a key exists. Panics on nil key.
	Has(key []byte) bool

	// Set sets the key. Panics on nil key or value.
	Set(key, value []byte)

	// Delete deletes the key. Panics on nil key.
	Delete(key []byte)
}
