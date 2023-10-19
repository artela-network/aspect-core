package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)

const (
	ARTELA_FROM_ADDR = "0x0000000000000000000000000000000000A27E14"
)

func IsArtelaFrom(from *common.Address) bool {
	if from != nil && from.String() == ARTELA_FROM_ADDR {
		return true
	}
	return false
}

func ArtelaAllowanceMsg(to *common.Address, nonce uint64, aspectId common.Address) (*core.Message, error) {
	from := common.HexToAddress(ARTELA_FROM_ADDR)
	return PackAllowanceMsg(from, to, nonce, big.NewInt(0), 9000000, big.NewInt(1000), big.NewInt(1000), big.NewInt(1000), nil, aspectId)
}
func ArtelaOwnerMsg(to *common.Address, nonce uint64, sender common.Address) (*core.Message, error) {
	from := common.HexToAddress(ARTELA_FROM_ADDR)
	return PackIsOwnerMsg(from, to, nonce, big.NewInt(0), 9000000, big.NewInt(1000), big.NewInt(30000), big.NewInt(2000), nil, sender)
}
