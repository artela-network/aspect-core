package run

import (
	"encoding/hex"
	"fmt"

	"github.com/artela-network/artelasdk/types"
	"github.com/artela-network/runtime"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"
)

const (
	// module of hostapis
	moduleHostApi = "hostapi"
	moduleAbis    = "abi"
	moduleUtils   = "utils"
	moduleDebug   = "debug"

	// namespace of hostapis
	nsHostApi = "__HostApi__"
	nsAbis    = "__Abi__"
	nsUtils   = "__Util__"
	nsDebug   = "__debug__"

	// entrance of api functions
	ApiEntrance           = "execute"
	CheckBlockLevel       = "isBlockLevel"
	CheckTransactionLevel = "isTransactionLevel"
)

// Register keeps the properity owned by current
type Register struct {
	aspectID    string
	collection  *runtime.HostAPIRegistry
	errCallback func(message string)
}

func NewRegister(aspectID string) *Register {
	return &Register{
		aspectID:   aspectID,
		collection: runtime.NewHostAPIRegistry(),
	}
}
func (r *Register) SetErrCallback(errfunc func(message string)) {
	r.errCallback = errfunc
}

// HostApis return the collection of aspect runtime host apis
func (r *Register) HostApis() *runtime.HostAPIRegistry {
	r.registerApis(moduleHostApi, nsHostApi, r.apis())
	r.registerApis(moduleAbis, nsAbis, r.abis())
	r.registerApis(moduleUtils, nsUtils, r.utils())
	r.registerApis(moduleDebug, nsDebug, r.debug())
	return r.collection
}

func (r *Register) registerApis(module, namespace string, apis interface{}) {
	for method, fn := range apis.(map[string]interface{}) {
		// Here we cannot make new variable function to call fn in it,
		// and to pass it into AddApi in loop instead pass fn directly.
		r.collection.AddApi(runtime.Module(module), runtime.Namesapce(namespace), runtime.MethodName(method), fn)
	}
}

func (r *Register) apis() interface{} {
	return map[string]interface{}{
		"lastBlock": func() []byte {
			if types.GetHostApiHook == nil {
				data, _ := types.NewBlockRet(false, "host functions is not init", nil).MarshalProto()
				return data
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				data, _ := types.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			block, err := host.LastBlock()
			if err != nil {
				data, _ := types.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			data, _ := types.NewBlockRet(true, "", block).MarshalProto()
			return data
		},
		"currentBlock": func() []byte {
			if types.GetHostApiHook == nil {
				data, _ := types.NewBlockRet(false, "host functions is not init", nil).MarshalProto()
				return data
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				data, _ := types.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			block, err := host.CurrentBlock()
			if err != nil {
				data, _ := types.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			data, _ := types.NewBlockRet(true, "", block).MarshalProto()
			return data
		},
		"localCall": func(arg []byte) []byte {
			// TODO
			return nil
		},
		"getContext": func(key string) string {
			if types.GetHostApiHook == nil {
				return "host functions is not valid"
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return "host functions is not init"
			}
			value, err := host.GetContext(r.aspectID, key)
			if err != nil {
				return err.Error()
			}
			return value
		},
		"setContext": func(key string, value string) bool {
			if types.GetHostApiHook == nil {
				return false
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return false
			}
			err = host.SetContext(r.aspectID, key, value)
			if err != nil {
				return false
			}
			return true
		},
		"setAspectState": func(key string, value string) bool {
			if types.GetHostApiHook == nil {
				return false
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return false
			}
			setErr := host.SetAspectState(r.aspectID, key, value)
			if setErr != nil {
				return false
			}
			return true
		},
		"getAspectState": func(key string) string {
			if types.GetHostApiHook == nil {
				return "host functions is not valid"
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return "host functions is not init"
			}
			value, err := host.GetAspectState(r.aspectID, key)
			if err != nil {
				return err.Error()
			}
			return value
		},
		"getProperty": func(key string) string {
			if types.GetHostApiHook == nil {
				return "host functions is not valid"
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return "host functions is not init"
			}
			value, err := host.GetProperty(r.aspectID, key)
			if err != nil {
				return err.Error()
			}
			return value
		},
		"scheduleTx": func(arg []byte) bool {
			if types.GetHostApiHook == nil {
				return false
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return false
			}
			sch := &types.Schedule{}
			if err := proto.Unmarshal(arg, sch); err != nil {
				return false
			}
			sch.Id.AspectId = r.aspectID
			return host.ScheduleTx(sch)
		},
		"getStateChanges": func(addr string, variable string, key []byte) []byte {
			if types.GetHostApiHook == nil {
				return nil
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return nil
			}
			changes := host.GetStateChanges(addr, variable, key)
			if changes == nil {
				return nil
			}

			data, err := proto.Marshal(changes)
			if err != nil {
				return nil
			}
			return data
		},
		"hash": func(hasher int32, data []byte) []byte {
			fmt.Println(string(data))
			hashFunc, ok := hashers[Hasher(hasher)]
			if !ok {
				return nil
			}

			return hashFunc(data)
		},
		// receive account address, and return the balance of the account with hex format.
		"currentBalance": func(addr string) string {
			if types.GetHostApiHook == nil {
				return ""
			}
			host, err := types.GetHostApiHook()
			if err != nil {
				return ""
			}
			acct := common.HexToAddress(addr)
			balance, err := host.CurrentBalance(acct)
			if err != nil {
				return ""
			}

			return balance.Text(16)
		},
	}
}

func (r *Register) abis() interface{} {
	return map[string]interface{}{
		"decodeParams": func(t string, data []byte) []byte {
			decodes, err := decodeParams(t, data)
			if err != nil {
				return []byte{}
			}

			values := &types.Values{
				All: make([]*types.Value, len(decodes)),
			}
			for i, decoded := range decodes {
				typeValue := &types.TypeValue{}
				typeValue.SetValue(decoded)
				values.All[i] = typeValue.Value()
			}

			byteArray, err := proto.Marshal(values)
			if err != nil {
				return []byte{}
			}
			return byteArray
		},
		"encodeParams": func(t string, valueData []byte) []byte {
			values := &types.Values{}
			if err := proto.Unmarshal(valueData, values); err != nil {
				return []byte{}
			}
			vals := make([]interface{}, len(values.All))
			for i, value := range values.All {
				typeValue := types.NewTypeValue(value)
				vals[i] = typeValue.GetValue()
			}
			data, err := encodeParams(t, vals...)
			if err != nil {
				return []byte{}
			}
			return data
		},
	}
}

func (r *Register) debug() interface{} {
	return map[string]interface{}{
		"log": func(s string) {
			fmt.Println(s)
		},
	}
}

func (r *Register) utils() interface{} {
	return map[string]interface{}{
		"fromHexString": func(s string) []byte {
			data, err := hex.DecodeString(s)
			if err != nil {
				return []byte{}
			}
			return data
		},
		"toHexString": func(data []byte) string {
			return hex.EncodeToString(data)
		},
		"revert": func(msg string) {
			if r.errCallback != nil {
				r.errCallback(msg)
			}
		},
	}
}
