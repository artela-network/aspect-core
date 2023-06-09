package run

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/artela-network/runtime"
	"google.golang.org/protobuf/proto"
)

const (
	// module of hostapis
	moduleHostApi = "host"
	moduleAbis    = "abi"

	// namespace of hostapis
	nsHostApi = "__HostApi__"
	nsAbis    = "__Abi__"

	// entrance of api functions
	ApiEntrance = "execute"
)

// Register keeps the properity owned by current
type Register struct {
	aspectID   string
	collection *runtime.HostAPIRegistry
}

func NewRegister(aspectID string) *Register {
	return &Register{
		aspectID:   aspectID,
		collection: runtime.NewHostAPIRegistry(),
	}
}

// HostApis return the collection of aspect runtime host apis
func (r *Register) HostApis() *runtime.HostAPIRegistry {
	r.registerApis(moduleHostApi, nsHostApi, r.apis())
	r.registerApis(moduleAbis, nsAbis, r.abis())
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
	}
}

func (r *Register) abis() interface{} {
	return map[string]interface{}{
		"decode": func(t string, data []byte) []byte {
			// TODO. use decode types
			values := &types.Values{
				All: make([]*types.Value, 0),
			}
			val, err := decodeType(t, data)
			if err != nil {
				return []byte{}
			}
			typeValue := &TypeValue{}
			typeValue.SetValue(val)
			values.All = append(values.All, typeValue.value)

			byteArray, err := proto.Marshal(values)
			if err != nil {
				return []byte{}
			}
			return byteArray
		},
		"encode": func(t string, valueData []byte) []byte {
			values := &types.Values{}
			if err := proto.Unmarshal(valueData, values); err != nil {
				return []byte{}
			}
			vals := make([]interface{}, len(values.All))
			for i, value := range values.All {
				typeValue := &TypeValue{value: value}
				vals[i] = typeValue.GetValue()
			}
			data, err := encodeTypes(t, vals...)
			if err != nil {
				return []byte{}
			}
			return data
		},
	}
}
