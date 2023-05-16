package run

import (
	"github.com/artela-network/runtime"

	aspectType "github.com/artela-network/artelasdk/types"
)

const (
	// module of hostapis
	moduleHostApi = "lib"

	// namespace of hostapis
	nsHostApi = "__HostApi__"

	// entrance of api functions
	ApiEntrance = "execute"
)

var (
	apis = map[string]interface{}{
		"lastBlock": func() []byte {
			host, err := aspectType.GetHostApiHook()
			if err != nil {
				data, _ := aspectType.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			block, err := host.LastBlock()
			if err != nil {
				data, _ := aspectType.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			data, _ := aspectType.NewBlockRet(true, "", block).MarshalProto()
			return data
		},
		"currentBlock": func() []byte {
			host, err := aspectType.GetHostApiHook()
			if err != nil {
				data, _ := aspectType.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			block, err := host.CurrentBlock()
			if err != nil {
				data, _ := aspectType.NewBlockRet(false, err.Error(), nil).MarshalProto()
				return data
			}
			data, _ := aspectType.NewBlockRet(true, "", block).MarshalProto()
			return data
		},
		"localCall": func(arg []byte) []byte {
			// TODO
			return nil
		},
	}
)

// HostApis return the collection of aspect runtime host apis
func HostApis() *runtime.HostAPICollection {
	return hostApis(moduleHostApi, nsHostApi)
}

func hostApis(module, namespace string) *runtime.HostAPICollection {
	collection := runtime.NewHostAPICollection()

	for method, fn := range apis {
		// Here we cannot make new variable function to call fn in it,
		// and to pass it into AddApi in loop instead pass fn directly.
		collection.AddApi(module, namespace, method, fn)
	}
	return collection
}
