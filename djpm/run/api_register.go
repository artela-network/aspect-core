package run

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/artela-network/runtime"
	"google.golang.org/protobuf/proto"
)

const (
	// module of hostapis
	moduleHostApi = "lib"

	// namespace of hostapis
	nsHostApi = "__HostApi__"

	// entrance of api functions
	ApiEntrance = "execute"
)

// Register keeps the properity owned by current
type Register struct {
	aspectID string
}

func NewRegister(aspectID string) *Register {
	return &Register{aspectID: aspectID}
}

// HostApis return the collection of aspect runtime host apis
func (r *Register) HostApis() *runtime.HostAPIRegistry {
	return r.hostApis(moduleHostApi, nsHostApi)
}

func (r *Register) hostApis(module, namespace string) *runtime.HostAPIRegistry {
	collection := runtime.NewHostAPIRegistry()

	for method, fn := range r.apis().(map[string]interface{}) {
		// Here we cannot make new variable function to call fn in it,
		// and to pass it into AddApi in loop instead pass fn directly.
		collection.AddApi(runtime.Module(module), runtime.Namesapce(namespace), runtime.MethodName(method), fn)
	}
	return collection
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
		"getStateChanges": func(addr string, variable string, key string) []byte {
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
