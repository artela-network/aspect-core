package types

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
)

var aspectCoreAddr = common.HexToAddress("0x0000000000000000000000000000000000A27E14")

func IsAspectContractAddr(to *common.Address) bool {
	return to != nil && bytes.Equal(aspectCoreAddr.Bytes(), to.Bytes())
}
