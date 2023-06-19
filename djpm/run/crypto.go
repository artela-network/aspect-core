package run

import "github.com/ethereum/go-ethereum/crypto"

type Hasher uint8

const (
	keccak Hasher = iota
)

var hashers = map[Hasher]func(...[]byte) []byte{
	keccak: crypto.Keccak256,
}
