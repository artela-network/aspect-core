package run

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/require"

	aspectType "github.com/artela-network/aspect-core/types"
)

// Run "scripts/build-wasm.sh" in project root, before run this test.
func TestJoinPoint(t *testing.T) {
	logger := &aspectType.NoOpsLogger{}
	aspectType.InitRuntimePool(context.Background(), logger, 0, 0)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			cwd, _ := os.Getwd()
			raw, _ := os.ReadFile(path.Join(cwd, "./tests/wasm/release.wasm"))

			name := aspectType.PRE_CONTRACT_CALL_METHOD
			input := &aspectType.PreContractCallInput{}
			reqData, err := hexutil.Decode("0x0a5b0a1400000000000000000000000000000000000000001214066e91dfc5bcc92eb992dca2307fd373f4d6adbe1800222470a08231000000000000000000000000e2af7c239b4f2800a2f742d406628b4fc4b8a0d42a003094c9f40b120408e8d507")
			if err != nil {
				panic(err)
			}
			err = proto.Unmarshal(reqData, input)
			if err != nil {
				panic(err)
			}
			runner, err := NewRunner(context.Background(), logger, "0x5f61973A8cDdCc531a663f15A2a65A2781fa6D1c", 1, raw, false)
			require.Equal(t, nil, err)
			address := common.HexToAddress("0x066e91dfc5bcc92eb992dca2307fd373f4d6adbe")

			_, gas, err := runner.JoinPoint(name, 24978580, 125691, address, input)
			runner.Return()
			require.Equal(t, nil, err)
			require.Equal(t, uint64(24978580), gas)
			wg.Done()
		}()
	}
	wg.Wait()

	cwd, _ := os.Getwd()
	raw, _ := os.ReadFile(path.Join(cwd, "./tests/wasm/release.wasm"))

	name := aspectType.PRE_CONTRACT_CALL_METHOD
	input := &aspectType.PreContractCallInput{}
	reqData, err := hexutil.Decode("0x0a5b0a1400000000000000000000000000000000000000001214066e91dfc5bcc92eb992dca2307fd373f4d6adbe1800222470a08231000000000000000000000000e2af7c239b4f2800a2f742d406628b4fc4b8a0d42a003094c9f40b120408e8d507")
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(reqData, input)
	if err != nil {
		panic(err)
	}
	runner, err := NewRunner(context.Background(), logger, "0x5f61973A8cDdCc531a663f15A2a65A2781fa6D1c", 1, raw, true)
	require.Equal(t, nil, err)
	address := common.HexToAddress("0x066e91dfc5bcc92eb992dca2307fd373f4d6adbe")

	_, gas, err := runner.JoinPoint(name, 24978580, 125691, address, input)
	runner.Return()
	require.Equal(t, nil, err)
	require.Equal(t, uint64(24978580), gas)
	fmt.Println("all test completed!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

// Run "scripts/build-wasm.sh" in project root, before run this test.
// func TestIsOwner(t *testing.T) {
// 	// cwd, _ := os.Getwd()
// 	raw, err := os.ReadFile("/Users/admin/mytech/go-work/src/github.com/artela-network/aspect-example/new_test/build/release.wasm")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	runner, err := NewRunner(context.Background(), "", raw)
// 	require.Equal(t, nil, err)
// 	address := common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")

// 	ret, err := runner.IsOwner(99, 99, &address, "hello")
// 	require.Equal(t, nil, err)
// 	defer runner.Return()

// 	require.Equal(t, true, ret)
// }
