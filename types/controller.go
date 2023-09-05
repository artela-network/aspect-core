package types

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
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

type AspectDataAccess string

const (
	BLOCK_HEIGHT_ACCESS             AspectDataAccess = "$.block.height"
	BLOCK_HEADER_ACCESS             AspectDataAccess = "$.block.header"
	BLOCK_VOTEINFO_ACCESS           AspectDataAccess = "$.block.voteInfo"
	CURRENT_TX_TRANSACTION_ACCESS   AspectDataAccess = "$.current_tx.transaction"
	CURRENT_TX_STATE_CHANGES_ACCESS AspectDataAccess = "$.current_tx.state_changes"
	CURRENT_TX_CALL_STACKS_ACCESS   AspectDataAccess = "$.current_tx.call_stacks"
	CURRENT_TX_RECEIPT_ACCESS       AspectDataAccess = "$.current_tx.receipt"
	CURRENT_TX_CONTEXT_ACCESS       AspectDataAccess = "$.current_tx.context"
	ASPECT_STATE_ACCESS             AspectDataAccess = "$.aspect.state"
	ASPECT_PROPERTIES_ACCESS        AspectDataAccess = "$.aspect.properties"
	EVM_INTERMEDIATE_STATE_ACCESS   AspectDataAccess = "$.evm.intermediate_state"
	EVM_BLOCK_STATE_ACCESS          AspectDataAccess = "$.evm.block_state"
	EVM_ENVCONFIG_ACCESS            AspectDataAccess = "$.evm.envconfig"
)
