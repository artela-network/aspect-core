package api

import (
	"crypto"
	"crypto/sha256"
	types2 "github.com/artela-network/aspect-runtime/types"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func (r *Registry) cryptoAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"sha256": {
			Func: func(input []byte) ([]byte, error) {
				h := sha256.Sum256(input)
				return h[:], nil
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"ripemd160": {
			Func: func(input []byte) ([]byte, error) {
				hash := crypto.RIPEMD160.New()
				hash.Write(input)
				return common.LeftPadBytes(hash.Sum(nil), 32), nil
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"keccak": {
			Func: func(input []byte) ([]byte, error) {
				return ethcrypto.Keccak256(input), nil
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"ecRecover": {
			Func: func(input []byte) ([]byte, error) {
				const ecRecoverInputLength = 128

				input = common.RightPadBytes(input, ecRecoverInputLength)
				// "input" is (hash, v, r, s), each 32 bytes
				// but for ecrecover we want (r, s, v)

				r := new(big.Int).SetBytes(input[64:96])
				s := new(big.Int).SetBytes(input[96:128])
				v := input[63] - 27

				// tighter sig s values input homestead only apply to tx sigs
				if !allZero(input[32:63]) || !ethcrypto.ValidateSignatureValues(v, r, s, false) {
					return nil, nil
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
					return nil, nil
				}

				// the first byte of pubkey is bitcoin heritage
				return common.LeftPadBytes(ethcrypto.Keccak256(pubKey[1:])[12:], 32), nil
			},
			GasRule: types2.NewStaticGasRule(1),
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
