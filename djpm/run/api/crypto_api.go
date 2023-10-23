package api

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"math/big"

	"github.com/cosmos/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func (r *Register) cryptoApis() interface{} {
	return map[string]interface{}{
		"sha256": func(input []byte) []byte {
			h := sha256.Sum256(input)
			return h[:]
		},
		"base64Encode": func(input []byte) string {
			return base64.StdEncoding.EncodeToString(input)
		},
		"base64Decode": func(input string) []byte {
			bytes, _ := base64.StdEncoding.DecodeString(input)
			return bytes
		},
		"base58Encode": func(input []byte) string {
			return base58.Encode(input)
		},
		"base58Decode": func(input string) []byte {
			return base58.Decode(input)
		},
		"ripemd160": func(input []byte) []byte {
			hash := crypto.RIPEMD160.New()
			hash.Write(input)
			return common.LeftPadBytes(hash.Sum(nil), 32)
		},
		"keccak": func(input []byte) []byte {
			return ethcrypto.Keccak256(input)
		},
		"ecRecover": func(input []byte) []byte {
			const ecRecoverInputLength = 128

			input = common.RightPadBytes(input, ecRecoverInputLength)
			// "input" is (hash, v, r, s), each 32 bytes
			// but for ecrecover we want (r, s, v)

			r := new(big.Int).SetBytes(input[64:96])
			s := new(big.Int).SetBytes(input[96:128])
			v := input[63] - 27

			// tighter sig s values input homestead only apply to tx sigs
			if !allZero(input[32:63]) || !ethcrypto.ValidateSignatureValues(v, r, s, false) {
				return nil
			}
			// We must make sure not to modify the 'input', so placing the 'v' along with
			// the signature needs to be done on a new allocation
			sig := make([]byte, 65)
			copy(sig, input[64:128])
			sig[64] = v
			// v needs to be at the end for libsecp256k1
			pubKey, err := ethcrypto.Ecrecover(input[:32], sig)
			// make sure the public key is a valid one
			if err != nil {
				return nil
			}

			// the first byte of pubkey is bitcoin heritage
			return common.LeftPadBytes(ethcrypto.Keccak256(pubKey[1:])[12:], 32)
		},
	}
}

func allZero(b []byte) bool {
	for _, one := range b {
		if one != 0 {
			return false
		}
	}
	return true
}
