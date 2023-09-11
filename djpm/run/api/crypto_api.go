package api

import "github.com/ethereum/go-ethereum/crypto"

type Hasher uint8

const (
	keccak Hasher = iota
)

var hashers = map[Hasher]func(...[]byte) []byte{
	keccak: crypto.Keccak256,
}

func (r *Register) cryptoApis() interface{} {
	return map[string]interface{}{
		"hash": func(hasher int32, data []byte) []byte {
			// fmt.Println(string(data))
			hashFunc, ok := hashers[Hasher(hasher)]
			if !ok {
				return nil
			}

			return hashFunc(data)
		},
	}
}
