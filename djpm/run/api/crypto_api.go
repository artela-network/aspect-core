package api

import (
	"crypto"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"math/big"

	"github.com/artela-network/aspect-core/types"
	types2 "github.com/artela-network/aspect-runtime/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/blake2b"
	"google.golang.org/protobuf/proto"
)

var (
	bn256AddAddress       = common.BytesToAddress([]byte{6})
	bn256ScalarMulAddress = common.BytesToAddress([]byte{7})
	bn256PairingAddress   = common.BytesToAddress([]byte{8})
)

func (r *Registry) cryptoAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"sha256": {
			Func: func(input []byte) ([]byte, error) {
				h := sha256.Sum256(input)
				return h[:], nil
			},
			GasRule: types2.NewDynamicGasRule(6000, 7500),
		},
		"ripemd160": {
			Func: func(input []byte) ([]byte, error) {
				hash := crypto.RIPEMD160.New()
				hash.Write(input)
				return common.LeftPadBytes(hash.Sum(nil), 32), nil
			},
			GasRule: types2.NewDynamicGasRule(60000, 75000),
		},
		"keccak": {
			Func: func(input []byte) ([]byte, error) {
				return ethcrypto.Keccak256(input), nil
			},
			GasRule: types2.NewDynamicGasRule(3000, 3750),
		},
		"ecRecover": {
			Func: func(input []byte) ([]byte, error) {
				ecRecover := &types.EcRecoverInput{}
				err := proto.Unmarshal(input, ecRecover)
				if err != nil {
					return nil, err
				}

				// "input" is (hash, v, r, s), each 32 bytes
				// but for ecrecover we want (r, s, v)

				r := new(big.Int).SetBytes(ecRecover.R)
				s := new(big.Int).SetBytes(ecRecover.S)
				v := ecRecover.V[len(ecRecover.V)-1] - 27

				// tighter sig s values input homestead only apply to tx sigs
				if !allZero(ecRecover.V[:len(ecRecover.V)-1]) || !ethcrypto.ValidateSignatureValues(v, r, s, false) {
					return nil, nil
				}
				// We must make sure not to modify the 'input', so placing the 'v' along with
				// the signature needs to be done on a new allocation
				sig := make([]byte, 65)
				// copy(sig, input[64:128])
				copy(sig, ecRecover.R)
				copy(sig[32:], ecRecover.S)
				sig[64] = v
				// v needs to be at the end for libsecp256k1
				pubKey, err := ethcrypto.Ecrecover(ecRecover.Hash, sig)
				// make sure the public key is a valid one
				if err != nil {
					return nil, nil
				}

				// the first byte of pubkey is bitcoin heritage
				return common.LeftPadBytes(ethcrypto.Keccak256(pubKey[1:])[12:], 32), nil
			},
			GasRule: types2.NewStaticGasRule(3000),
		},
		"bigModExp": {
			// proto encoded input ([]byte) -> proto decode (3 []byte)
			Func: func(b, e, m []byte) ([]byte, error) {
				// Handle a special case when both the base and mod length is zero
				if len(b) == 0 && len(m) == 0 {
					return []byte{}, errors.New("params not valid")
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
					return common.LeftPadBytes([]byte{}, len(m)), nil
				case base.BitLen() == 1: // a bit length of 1 means it's 1 (or -1).
					// If base == 1, then we can just return base % mod (if mod >= 1, which it is)
					v = base.Mod(base, mod).Bytes()
				default:
					v = base.Exp(base, exp, mod).Bytes()
				}
				return common.LeftPadBytes(v, len(m)), nil
			},
			GasRule: types2.NewDynamicGasRule(15000, 300000),
		},

		"bn256Add": {
			Func: func(input []byte) ([]byte, error) {
				points := &types.Bn256AddInput{}
				err := proto.Unmarshal(input, points)
				if err != nil {
					return nil, err
				}

				calldata := make([]byte, 128)
				copy(calldata[:], points.A.X)
				copy(calldata[32:], points.A.Y)
				copy(calldata[64:], points.B.X)
				copy(calldata[96:], points.B.Y)

				contract := vm.PrecompiledContractsIstanbul[bn256AddAddress]
				res, err := contract.Run(calldata)
				if err != nil {
					return nil, err
				}

				if len(res) != 64 {
					return nil, errors.New("run precompile failed")
				}
				point := &types.G1{X: res[:32], Y: res[32:]}
				return proto.Marshal(point)
			},
			GasRule: types2.NewStaticGasRule(1500),
		},

		"bn256ScalarMul": {
			Func: func(input []byte) ([]byte, error) {
				scalrInput := &types.Bn256ScalarMulInput{}
				err := proto.Unmarshal(input, scalrInput)
				if err != nil {
					return nil, err
				}

				calldata := make([]byte, 96)
				copy(calldata[:], scalrInput.A.X)
				copy(calldata[32:], scalrInput.A.Y)
				copy(calldata[64:], scalrInput.Scalar)

				contract := vm.PrecompiledContractsIstanbul[bn256ScalarMulAddress]
				res, err := contract.Run(calldata)
				if err != nil {
					return nil, err
				}

				if len(res) != 64 {
					return nil, errors.New("run precompile failed")
				}
				scalared := &types.G1{X: res[:32], Y: res[32:]}
				return proto.Marshal(scalared)
			},
			GasRule: types2.NewStaticGasRule(6000),
		},

		"bn256Pairing": {
			Func: func(input []byte) ([]byte, error) {
				pairing := &types.Bn256PairingInput{}
				err := proto.Unmarshal(input, pairing)
				if err != nil {
					return nil, err
				}

				if !(len(pairing.Cs) == len(pairing.Ts)) {
					return nil, errors.New("params not valid")
				}

				grouplen := 192 // 32 * 2 * 3

				calldata := make([]byte, len(pairing.Cs)*grouplen)
				for i := 0; i < len(pairing.Cs); i++ {
					start := grouplen * i
					copy(calldata[start:], pairing.Cs[i].X)
					copy(calldata[start+32:], pairing.Cs[i].Y)
					copy(calldata[start+64:], pairing.Ts[i].X1)
					copy(calldata[start+96:], pairing.Ts[i].X2)
					copy(calldata[start+128:], pairing.Ts[i].Y1)
					copy(calldata[start+160:], pairing.Ts[i].Y2)
				}

				contract := vm.PrecompiledContractsIstanbul[bn256PairingAddress]
				res, err := contract.Run(calldata)
				if err != nil {
					return nil, err
				}

				return res, nil
			},
			GasRule: types2.NewDynamicGasRule(45000, 380000),
		},

		"blake2F": {
			Func: func(input []byte) ([]byte, error) {
				// Execute the compression function, extract and return the result
				blakeInput := &types.Blake2FInput{}
				err := proto.Unmarshal(input, blakeInput)
				if err != nil {
					return nil, err
				}

				if !(len(blakeInput.H) == 64 && len(blakeInput.M) == 128 && len(blakeInput.T) == 16 && blakeInput.Final != nil && blakeInput.Rounds != nil) {
					return nil, errors.New("params not valid")
				}

				var (
					h [8]uint64
					m [16]uint64
					t [2]uint64

					rounds = binary.BigEndian.Uint32(blakeInput.Rounds[0:4])
				)

				for i := 0; i < 8; i++ {
					offset := i*8
					h[i] = binary.LittleEndian.Uint64(blakeInput.H[offset : offset+8])
				}
				for i := 0; i < 16; i++ {
					offset := i*8
					m[i] = binary.LittleEndian.Uint64(blakeInput.M[offset : offset+8])
				}
				t[0] = binary.LittleEndian.Uint64(blakeInput.T[:8])
				t[1] = binary.LittleEndian.Uint64(blakeInput.T[8:16])

				blake2b.F(&h, m, t, *blakeInput.Final, rounds)

				output := make([]byte, 64)
				for i := 0; i < 8; i++ {
					offset := i * 8
					binary.LittleEndian.PutUint64(output[offset:offset+8], h[i])
				}
				return output, nil
			},
			GasRule: types2.NewDynamicGasRule(30000, 150000),
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
