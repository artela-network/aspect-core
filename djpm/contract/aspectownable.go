package contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// 1. brew tap ethereum/ethereum
// 2. brew install solidity
//go:generate solc onwer.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ./ --overwrite
//go:generate abigen --pkg contract --out aspectownable_contract.go --combined-json ./combined.json

var AspectOwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2f54bf6e": "isOwner(address)",
	},
}

func PackIsOwnerMsg(from common.Address, to *common.Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, accessList types.AccessList, sender common.Address) (core.Message, error) {
	parsed, err := AspectOwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	// Pack the input, call and unpack the results
	input, err := parsed.Pack("isOwner", &sender)
	if err != nil {
		return nil, err
	}
	message := types.NewMessage(from, to, nonce, amount, gasLimit, gasPrice, gasFeeCap, gasTipCap, input, accessList, false)
	return message, nil
}

func UnpackIsOwnerResult(data []byte) (bool, error) {
	parsed, err := AspectOwnableMetaData.GetAbi()
	if err != nil {
		return false, err
	}
	maps := make(map[string]interface{}, 0)
	err2 := parsed.UnpackIntoMap(maps, "isOwner", data)
	if err2 != nil {
		fmt.Println("pack error", err2)
	}
	out0 := *abi.ConvertType(maps["result"], new(bool)).(*bool)
	return out0, nil
}
