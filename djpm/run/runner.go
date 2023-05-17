package run

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/artela-network/runtime/wasmtime"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func RunAspect(code []byte, method string, input *types.AspectInput) (*types.AspectOutput, error) {
	wasmTimeRuntime, err := wasmtime.NewWASMTimeRuntime(code, HostApis())
	if err != nil {
		return nil, err
	}

	// turn input into bytes
	reqData, err := proto.Marshal(input)
	if err != nil {
		return nil, err
	}
	res, err := wasmTimeRuntime.Call(ApiEntrance, method, reqData)
	if err != nil {
		return nil, err
	}

	resData, ok := res.([]byte)
	if !ok {
		return nil, errors.New("read output failed, return value is not byte array")
	}

	output := &types.AspectOutput{}
	if err := proto.Unmarshal(resData, output); err != nil {
		return nil, errors.Wrap(err, "unmarshal AspectOutput")
	}

	return output, nil
}
