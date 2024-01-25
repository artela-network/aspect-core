package api

import (
	"github.com/artela-network/aspect-core/types"
	"google.golang.org/protobuf/proto"
)

func (r *Registry) blockchainAPIs() interface{} {
	return map[string]interface{}{
		"getTransactionByHash": func(hash []byte) []byte {
			hook, err := types.GetBlockchainHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetBlockchainHook failed")
			}
			tx := hook.GetTransactionByHash(hash)
			marshal, err := proto.Marshal(tx)
			if err != nil {
				// not get the correct transaction, send a
				empty := uint64(0)
				marshal, _ = proto.Marshal(&types.Transaction{
					BlockHash:   []byte{},
					BlockNumber: &empty,
					Hash:        hash,
				})
			}
			return wrapNilByte(marshal)
		},
	}
}
