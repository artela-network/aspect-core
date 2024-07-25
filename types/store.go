package types

type AspectStore interface {
	Get(key []byte) []byte

	// Has checks if a key exists. Panics on nil key.
	Has(key []byte) bool

	// Set sets the key. Panics on nil key or value.
	Set(key, value []byte)

	// Delete deletes the key. Panics on nil key.
	Delete(key []byte)
}
