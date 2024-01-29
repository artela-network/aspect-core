package api

import (
	"crypto"
	"crypto/sha256"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func (r *Registry) cryptoAPIs() interface{} {
	return map[string]interface{}{
		"sha256": func(input []byte) []byte {
			h := sha256.Sum256(input)
			return h[:]
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
		"bigModExp": func(b, e, m []byte) []byte {
			if len(b) > 32 || len(e) > 32 || len(m) > 32 {
				return []byte{}
			}

			var (
				baseLen = new(big.Int).SetBytes(b).Uint64()
				// expLen  = new(big.Int).SetBytes(e).Uint64()
				modLen = new(big.Int).SetBytes(m).Uint64()
			)
			// Handle a special case when both the base and mod length is zero
			if baseLen == 0 && modLen == 0 {
				return []byte{}
			}
			// Retrieve the operands and execute the exponentiation
			var (
				base = new(big.Int).SetBytes(b)
				exp  = new(big.Int).SetBytes(e)
				mod  = new(big.Int).SetBytes(m)
				v    []byte
			)
			switch {
			case mod.BitLen() == 0:
				// Modulo 0 is undefined, return zero
				return common.LeftPadBytes([]byte{}, int(modLen))
			case base.BitLen() == 1: // a bit length of 1 means it's 1 (or -1).
				// If base == 1, then we can just return base % mod (if mod >= 1, which it is)
				v = base.Mod(base, mod).Bytes()
			default:
				v = base.Exp(base, exp, mod).Bytes()
			}
			return common.LeftPadBytes(v, int(modLen))
		},

		"bn256Add": func(ax, ay, bx, by []byte) []byte {
			if len(ax) > 32 || len(ay) > 32 || len(bx) > 32 || len(by) > 32 {
				return []byte{}
			}
			input := make([]byte, 128)
			copy(input[:], ax)
			copy(input[32:], ay)
			copy(input[64:], bx)
			copy(input[96:], by)

			bn256AddAddress := common.BytesToAddress([]byte{6})
			contract := vm.PrecompiledContractsBerlin[bn256AddAddress]
			data, err := contract.Run(input)
			if err != nil {
				return nil
			}
			return data
		},

		"bn256ScalarMul": func(x, y, scalar []byte) []byte {
			if len(x) > 32 || len(y) > 32 || len(scalar) > 32 {
				return []byte{}
			}
			input := make([]byte, 96)
			copy(input[:], x)
			copy(input[32:], y)
			copy(input[64:], scalar)

			bn256ScalarMulAddress := common.BytesToAddress([]byte{7})
			contract := vm.PrecompiledContractsBerlin[bn256ScalarMulAddress]
			data, err := contract.Run(input)
			if err != nil {
				return nil
			}
			return data
		},

		"bn256Pairing": func(input []byte) []byte {
			bn256AddAddress := common.BytesToAddress([]byte{6})
			contract := vm.PrecompiledContractsBerlin[bn256AddAddress]
			data, err := contract.Run(input)
			if err != nil {
				return nil
			}
			return data
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
