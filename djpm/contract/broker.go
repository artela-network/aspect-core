package contract

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// 1. brew tap ethereum/ethereum
// 2. brew install solidity
//go:generate solc broker.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ./ --overwrite
//go:generate abigen --pkg contract --out broker_contract.go --combined-json ./combined.json

var BrokMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aspectId\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"valueWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3e5beab9": "allowance(address)",
	},
}

func PackAllowanceMsg(from common.Address, to *common.Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, accessList types.AccessList, aspectId common.Address) (*core.Message, error) {
	parsed, err := BrokMetaData.GetAbi()
	if err != nil {
		return &core.Message{}, err
	}
	// Pack the input, call and unpack the results
	input, err := parsed.Pack("allowance", aspectId)
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

func UnpackAllowanceResult(data []byte) (*big.Int, error) {
	parsed, err := BrokMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	maps := make(map[string]interface{}, 0)
	err2 := parsed.UnpackIntoMap(maps, "allowance", data)
	if err2 != nil {
		fmt.Println("pack error", err2)
	}
	out0 := *abi.ConvertType(maps["valueWei"], new(*big.Int)).(**big.Int)
	return out0, nil
}
