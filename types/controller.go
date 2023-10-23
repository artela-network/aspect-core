package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ARTELA_ADDR = "0x0000000000000000000000000000000000A27E14"
)

func IsAspectContract(to string) bool {
	if to != "" && strings.EqualFold(ARTELA_ADDR, to) {
		// ignore contract deployment transaction & aspect op txs
		return true
	}
	return false
}

func IsAspectContractAddr(to *common.Address) bool {
	if to != nil && strings.EqualFold(ARTELA_ADDR, to.String()) {
		// ignore contract deployment transaction & aspect op txs
		return true
	}
	return false
}
