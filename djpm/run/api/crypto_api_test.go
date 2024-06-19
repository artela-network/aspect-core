package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/artela-network/aspect-core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

type precompiledTest struct {
	Input, Expected string
	Gas             uint64
	Name            string
	NoBenchmark     bool // Benchmark primarily the worst-cases
}

func loadJson(name string) ([]precompiledTest, error) {
	data, err := os.ReadFile(fmt.Sprintf("testdata/%v.json", name))
	if err != nil {
		return nil, err
	}
	var testcases []precompiledTest
	err = json.Unmarshal(data, &testcases)
	return testcases, err
}

func getData(data []byte, start uint64, size uint64) []byte {
	length := uint64(len(data))
	if start > length {
		start = length
	}
	end := start + size
	if end > length {
		end = length
	}
	return common.RightPadBytes(data[start:end], int(size))
}

func TestEcRecover(t *testing.T) {
	r := NewRegistry(context.Background(), common.Address{}, 1)
	apis := r.cryptoAPIs()

	api, ok := apis["ecRecover"]
	require.Equal(t, true, ok)

	fn, ok := api.Func.(func(input []byte) ([]byte, error))
	require.Equal(t, true, ok)

	testCases, err := loadJson("ecRecover")
	require.Equal(t, nil, err)
	for _, testCase := range testCases {
		input := common.Hex2Bytes(testCase.Input)

		hash := getData(input, 0, 32)
		v := getData(input, 32, 32)
		r := getData(input, 64, 32)
		s := getData(input, 96, 32)

		fmt.Println("hash: ", common.Bytes2Hex(hash))
		fmt.Println("v: ", common.Bytes2Hex(v))
		fmt.Println("r: ", common.Bytes2Hex(r))
		fmt.Println("s: ", common.Bytes2Hex(s))
		fmt.Println("expect: ", testCase.Expected)

		data, err := proto.Marshal(&types.EcRecoverInput{
			Hash: hash,
			V: v,
			R: r,
			S: s,
		})
		require.Equal(t, nil, err)

		res, err := fn(data)
		require.Equal(t, nil, err)

		require.Equal(t, testCase.Expected, common.Bytes2Hex(res))
	}
}

func TestModExp(t *testing.T) {
	r := NewRegistry(context.Background(), common.Address{}, 1)
	apis := r.cryptoAPIs()

	api, ok := apis["bigModExp"]
	require.Equal(t, true, ok)

	fn, ok := api.Func.(func(b, e, m []byte) ([]byte, error))
	require.Equal(t, true, ok)

	testCases, err := loadJson("modexp")
	require.Equal(t, nil, err)
	for _, testCase := range testCases {
		input, err := hex.DecodeString(testCase.Input)
		require.Equal(t, nil, err)

		baseLen := new(big.Int).SetBytes(getData(input, 0, 32)).Uint64()
		expLen := new(big.Int).SetBytes(getData(input, 32, 32)).Uint64()
		modLen := new(big.Int).SetBytes(getData(input, 64, 32)).Uint64()

		if len(input) > 96 {
			input = input[96:]
		} else {
			input = input[:0]
		}

		b := getData(input, 0, baseLen)
		e := getData(input, baseLen, expLen)
		m := getData(input, baseLen+expLen, modLen)

		res, err := fn(b, e, m)
		require.Equal(t, nil, err)
		require.Equal(t, testCase.Expected, hex.EncodeToString(res))
	}
}

func TestBN256Add(t *testing.T) {
	r := NewRegistry(context.Background(), common.Address{}, 1)
	apis := r.cryptoAPIs()

	api, ok := apis["bn256Add"]
	require.Equal(t, true, ok)

	fn, ok := api.Func.(func(input []byte) ([]byte, error))
	require.Equal(t, true, ok)

	testCases, err := loadJson("bn256Add")
	require.Equal(t, nil, err)
	for _, testCase := range testCases {
		input := common.Hex2Bytes(testCase.Input)

		ax := getData(input, 0, 32)
		ay := getData(input, 32, 32)
		bx := getData(input, 64, 32)
		by := getData(input, 96, 32)

		data, err := proto.Marshal(&types.Bn256AddInput{
			A: &types.G1{X: ax, Y: ay},
			B: &types.G1{X: bx, Y: by},
		})
		require.Equal(t, nil, err)

		res, err := fn(data)
		require.Equal(t, nil, err)

		p := &types.G1{}
		err = proto.Unmarshal(res, p)
		require.Equal(t, nil, err)

		require.Equal(t, testCase.Expected[:64], common.Bytes2Hex(p.X))
		require.Equal(t, testCase.Expected[64:], common.Bytes2Hex(p.Y))
	}
}

func TestBN256Saclar(t *testing.T) {
	r := NewRegistry(context.Background(), common.Address{}, 1)
	apis := r.cryptoAPIs()

	api, ok := apis["bn256ScalarMul"]
	require.Equal(t, true, ok)

	fn, ok := api.Func.(func(input []byte) ([]byte, error))
	require.Equal(t, true, ok)

	testCases, err := loadJson("bn256ScalarMul")
	require.Equal(t, nil, err)
	for _, testCase := range testCases {
		input := common.Hex2Bytes(testCase.Input)

		x := getData(input, 0, 32)
		y := getData(input, 32, 32)
		scalar := getData(input, 64, 32)

		scalarInput := &types.Bn256ScalarMulInput{A: &types.G1{X: x, Y: y}, Scalar: scalar}
		scalarData, err := proto.Marshal(scalarInput)
		require.Equal(t, nil, err)

		res, err := fn(scalarData)
		require.Equal(t, nil, err)

		p := &types.G1{}
		err = proto.Unmarshal(res, p)
		require.Equal(t, nil, err)

		require.Equal(t, testCase.Expected[:64], common.Bytes2Hex(p.X))
		require.Equal(t, testCase.Expected[64:], common.Bytes2Hex(p.Y))
	}
}

func TestBN256Pairing(t *testing.T) {
	r := NewRegistry(context.Background(), common.Address{}, 1)
	apis := r.cryptoAPIs()

	api, ok := apis["bn256Pairing"]
	require.Equal(t, true, ok)

	fn, ok := api.Func.(func(input []byte) ([]byte, error))
	require.Equal(t, true, ok)

	testCases, err := loadJson("bn256Pairing")
	require.Equal(t, nil, err)
	for _, testCase := range testCases {
		input := common.Hex2Bytes(testCase.Input)
		pairing := &types.Bn256PairingInput{}
		for i := 0; i < len(input); i += 192 {
			c, t := &types.G1{}, &types.G2{}
			c.X = input[i : i+32]
			c.Y = input[i+32 : i+(32*2)]
			t.X1 = input[i+(32*2) : i+(32*3)]
			t.X2 = input[i+(32*3) : i+(32*4)]
			t.Y1 = input[i+(32*4) : i+(32*5)]
			t.Y2 = input[i+(32*5) : i+(32*6)]
			pairing.Cs = append(pairing.Cs, c)
			pairing.Ts = append(pairing.Ts, t)
		}

		in, err := proto.Marshal(pairing)
		require.Equal(t, nil, err)
		res, err := fn(in)
		require.Equal(t, nil, err)
		require.Equal(t, 32, len(res))
		require.Equal(t, testCase.Expected, common.Bytes2Hex(res))
	}
}
