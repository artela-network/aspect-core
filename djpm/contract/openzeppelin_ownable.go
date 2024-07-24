package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// 1. brew tap ethereum/ethereum
// 2. brew install solidity
//go:generate solc openzeppelin_ownable.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ./ --overwrite
//go:generate abigen --pkg contract --out openzeppelin_ownable_contract.go --combined-json ./combined.json

var OpenZeppelinOwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
	},
}

func PackOwnableOwnerMsg(from common.Address, to *common.Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, accessList types.AccessList, _ common.Address) (*core.Message, error) {
	parsed, err := OpenZeppelinOwnableMetaData.GetAbi()
	if err != nil {
		return &core.Message{}, err
	}
	// Pack the input, call and unpack the results
	input, err := parsed.Pack("owner")
	if err != nil {
		return &core.Message{}, err
	}
	message := &core.Message{
		To:                to,
		From:              from,
		Nonce:             nonce,
		Value:             amount,
		GasLimit:          gasLimit,
		GasPrice:          gasPrice,
		GasFeeCap:         gasFeeCap,
		GasTipCap:         gasTipCap,
		Data:              input,
		AccessList:        accessList,
		SkipAccountChecks: false,
	}
	return message, nil
}

func UnpackOwnableOwnerResult(data []byte) (common.Address, error) {
	parsed, err := OpenZeppelinOwnableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}
	maps := make(map[string]interface{}, 1)
	if err = parsed.UnpackIntoMap(maps, "owner", data); err != nil {
		return common.Address{}, err
	}
	out0 := *abi.ConvertType(maps[""], new(common.Address)).(*common.Address)
	return out0, nil
}
