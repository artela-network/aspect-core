package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
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

func ArtelaOwnerMsg(to *common.Address, nonce uint64, sender common.Address, gas uint64, gasPrice, gasFeeCap, gasTipCap *big.Int) (*core.Message, error) {
	from := common.HexToAddress(ARTELA_FROM_ADDR)
	return PackIsOwnerMsg(from, to, nonce, big.NewInt(0), gas, gasPrice, gasFeeCap, gasTipCap, nil, sender)
}

func OpenZeppelinOwnableMsg(to *common.Address, nonce uint64, sender common.Address, gas uint64, gasPrice, gasFeeCap, gasTipCap *big.Int) (*core.Message, error) {
	from := common.HexToAddress(ARTELA_FROM_ADDR)
	return PackIsOwnerMsg(from, to, nonce, big.NewInt(0), gas, gasPrice, gasFeeCap, gasTipCap, nil, sender)
}
