package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MarshalJSON marshals as JSON.
func (x *NoFromTxInput) MarshalJSON() ([]byte, error) {
	type noFromTxInput struct {
		Hash common.Hash    `json:"hash,omitempty"`
		To   common.Address `json:"to,omitempty"`
	}
	var enc noFromTxInput
	enc.Hash = common.BytesToHash(x.Hash)
	enc.To = common.BytesToAddress(x.To)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *NoFromTxInput) UnmarshalJSON(input []byte) error {
	type noFromTxInput struct {
		Hash common.Hash    `json:"hash,omitempty"`
		To   common.Address `json:"to,omitempty"`
	}
	var dec noFromTxInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	x.Hash = dec.Hash.Bytes()
	x.To = dec.To.Bytes()
	return nil
}

// MarshalJSON marshals as JSON.
func (x *WithFromTxInput) MarshalJSON() ([]byte, error) {
	type withFromTxInput struct {
		Hash common.Hash    `json:"hash,omitempty"`
		To   common.Address `json:"to,omitempty"`
		From common.Address `json:"from,omitempty"`
	}
	var enc withFromTxInput
	enc.Hash = common.BytesToHash(x.Hash)
	enc.To = common.BytesToAddress(x.To)
	enc.From = common.BytesToAddress(x.From)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *WithFromTxInput) UnmarshalJSON(input []byte) error {
	type withFromTxInput struct {
		Hash common.Hash    `json:"hash,omitempty"`
		To   common.Address `json:"to,omitempty"`
		From common.Address `json:"from,omitempty"`
	}
	var dec withFromTxInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	x.Hash = dec.Hash.Bytes()
	x.To = dec.To.Bytes()
	x.From = dec.From.Bytes()
	return nil
}

// MarshalJSON marshals as JSON.
func (x *PreExecMessageInput) MarshalJSON() ([]byte, error) {
	type preExecMessageInput struct {
		From  common.Address `json:"from,omitempty"`
		To    common.Address `json:"to,omitempty"`
		Index hexutil.Uint64 `json:"index,omitempty"`
		Data  hexutil.Bytes  `json:"data,omitempty"`
		Value hexutil.Bytes  `json:"value,omitempty"`
		Gas   hexutil.Uint64 `json:"gas,omitempty"`
	}
	var enc preExecMessageInput
	enc.From = common.BytesToAddress(x.From)
	enc.To = common.BytesToAddress(x.To)
	if x.Index != nil {
		enc.Index = hexutil.Uint64(*x.Index)
	}
	enc.Data = x.Data
	enc.Value = x.Value
	if x.Gas != nil {
		enc.Gas = hexutil.Uint64(*x.Gas)
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *PreExecMessageInput) UnmarshalJSON(input []byte) error {
	type preExecMessageInput struct {
		From  common.Address  `json:"from,omitempty"`
		To    common.Address  `json:"to,omitempty"`
		Index *hexutil.Uint64 `json:"index,omitempty"`
		Data  hexutil.Bytes   `json:"data,omitempty"`
		Value hexutil.Bytes   `json:"value,omitempty"`
		Gas   *hexutil.Uint64 `json:"gas,omitempty"`
	}
	var dec preExecMessageInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	x.From = dec.From.Bytes()
	x.To = dec.To.Bytes()
	if dec.Index != nil {
		x.Index = (*uint64)(dec.Index)
	}
	x.Data = dec.Data
	x.Value = dec.Value
	if dec.Gas != nil {
		x.Gas = (*uint64)(dec.Gas)
	}
	return nil
}

// MarshalJSON marshals as JSON.
func (x *PostExecMessageInput) MarshalJSON() ([]byte, error) {
	type postExecMessageInput struct {
		From  common.Address `json:"from,omitempty"`
		To    common.Address `json:"to,omitempty"`
		Index hexutil.Uint64 `json:"index,omitempty"`
		Data  hexutil.Bytes  `json:"data,omitempty"`
		Value hexutil.Bytes  `json:"value,omitempty"`
		Gas   hexutil.Uint64 `json:"gas,omitempty"`
		Ret   hexutil.Bytes  `json:"ret,omitempty"`
		Error *string        `json:"error,omitempty"`
	}
	var enc postExecMessageInput
	enc.From = common.BytesToAddress(x.From)
	enc.To = common.BytesToAddress(x.To)
	if x.Index != nil {
		enc.Index = hexutil.Uint64(*x.Index)
	}
	enc.Data = x.Data
	enc.Value = x.Value
	if x.Gas != nil {
		enc.Gas = hexutil.Uint64(*x.Gas)
	}
	enc.Ret = x.Ret
	if x.Error != nil && *x.Error != "" {
		enc.Error = x.Error
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *PostExecMessageInput) UnmarshalJSON(input []byte) error {
	type postExecMessageInput struct {
		From  common.Address  `json:"from,omitempty"`
		To    common.Address  `json:"to,omitempty"`
		Index *hexutil.Uint64 `json:"index,omitempty"`
		Data  hexutil.Bytes   `json:"data,omitempty"`
		Value hexutil.Bytes   `json:"value,omitempty"`
		Gas   *hexutil.Uint64 `json:"gas,omitempty"`
		Ret   hexutil.Bytes   `json:"ret,omitempty"`
		Error *string         `json:"error,omitempty"`
	}
	var dec postExecMessageInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	x.From = dec.From.Bytes()
	x.To = dec.To.Bytes()
	if dec.Index != nil {
		x.Index = (*uint64)(dec.Index)
	}
	x.Data = dec.Data
	x.Value = dec.Value
	if dec.Gas != nil {
		x.Gas = (*uint64)(dec.Gas)
	}
	x.Ret = dec.Ret
	x.Error = dec.Error
	return nil
}

// MarshalJSON marshals as JSON.
func (x *BlockInput) MarshalJSON() ([]byte, error) {
	type blockInput struct {
		Number hexutil.Uint64 `json:"number,omitempty"`
	}
	var enc blockInput
	if x.Number != nil {
		enc.Number = hexutil.Uint64(*x.Number)
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *BlockInput) UnmarshalJSON(input []byte) error {
	type blockInput struct {
		Number *hexutil.Uint64 `json:"number,omitempty"`
	}
	var dec blockInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Number != nil {
		x.Number = (*uint64)(dec.Number)
	}
	return nil
}

// MarshalJSON marshals as JSON.
func (x *ReceiptInput) MarshalJSON() ([]byte, error) {
	type receiptInput struct {
		Status hexutil.Uint64 `json:"status,omitempty"`
	}
	var enc receiptInput
	if x.Status != nil {
		enc.Status = hexutil.Uint64(*x.Status)
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *ReceiptInput) UnmarshalJSON(input []byte) error {
	type receiptInput struct {
		Status *hexutil.Uint64 `json:"status,omitempty"`
	}
	var dec receiptInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Status != nil {
		x.Status = (*uint64)(dec.Status)
	}
	return nil
}

// MarshalJSON marshals as JSON.
func (x *TxVerifyInput) MarshalJSON() ([]byte, error) {
	type txVerifyInput struct {
		Tx             *NoFromTxInput `json:"tx,omitempty"`
		Block          *BlockInput    `json:"block,omitempty"`
		ValidationData hexutil.Bytes  `json:"validation_data,omitempty"`
		CallData       hexutil.Bytes  `json:"call_data,omitempty"`
	}
	var enc txVerifyInput
	if x.Tx != nil {
		enc.Tx = x.Tx
	}
	if x.Block != nil {
		enc.Block = x.Block
	}
	enc.ValidationData = x.ValidationData
	enc.CallData = x.CallData
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *TxVerifyInput) UnmarshalJSON(input []byte) error {
	type txVerifyInput struct {
		Tx             *NoFromTxInput `json:"tx,omitempty"`
		Block          *BlockInput    `json:"block,omitempty"`
		ValidationData hexutil.Bytes  `json:"validation_data,omitempty"`
		CallData       hexutil.Bytes  `json:"call_data,omitempty"`
	}
	var dec txVerifyInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Tx != nil {
		x.Tx = dec.Tx
	}
	if dec.Block != nil {
		x.Block = dec.Block
	}
	x.ValidationData = dec.ValidationData
	x.CallData = dec.CallData
	return nil
}

// MarshalJSON marshals as JSON.
func (x *PreTxExecuteInput) MarshalJSON() ([]byte, error) {
	type preTxExecuteInput struct {
		Tx    *WithFromTxInput `json:"tx,omitempty"`
		Block *BlockInput      `json:"block,omitempty"`
	}
	var enc preTxExecuteInput
	if x.Tx != nil {
		enc.Tx = x.Tx
	}
	if x.Block != nil {
		enc.Block = x.Block
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *PreTxExecuteInput) UnmarshalJSON(input []byte) error {
	type preTxExecuteInput struct {
		Tx    *WithFromTxInput `json:"tx,omitempty"`
		Block *BlockInput      `json:"block,omitempty"`
	}
	var dec preTxExecuteInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Tx != nil {
		x.Tx = dec.Tx
	}
	if dec.Block != nil {
		x.Block = dec.Block
	}
	return nil
}

// MarshalJSON marshals as JSON.
func (x *OperationInput) MarshalJSON() ([]byte, error) {
	type operationInput struct {
		Tx       *WithFromTxInput `json:"tx,omitempty"`
		Block    *BlockInput      `json:"block,omitempty"`
		CallData hexutil.Bytes    `json:"call_data,omitempty"`
	}
	var enc operationInput
	if x.Tx != nil {
		enc.Tx = x.Tx
	}
	if x.Block != nil {
		enc.Block = x.Block
	}
	enc.CallData = x.CallData
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *OperationInput) UnmarshalJSON(input []byte) error {
	type operationInput struct {
		Tx       *WithFromTxInput `json:"tx,omitempty"`
		Block    *BlockInput      `json:"block,omitempty"`
		CallData hexutil.Bytes    `json:"call_data,omitempty"`
	}
	var dec operationInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Tx != nil {
		x.Tx = dec.Tx
	}
	if dec.Block != nil {
		x.Block = dec.Block
	}
	x.CallData = dec.CallData
	return nil
}

// MarshalJSON marshals as JSON.
func (x *InitInput) MarshalJSON() ([]byte, error) {
	type initInput struct {
		Tx       *WithFromTxInput `json:"tx,omitempty"`
		Block    *BlockInput      `json:"block,omitempty"`
		CallData hexutil.Bytes    `json:"call_data,omitempty"`
	}
	var enc initInput
	if x.Tx != nil {
		enc.Tx = x.Tx
	}
	if x.Block != nil {
		enc.Block = x.Block
	}
	enc.CallData = x.CallData
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *InitInput) UnmarshalJSON(input []byte) error {
	type initInput struct {
		Tx       *WithFromTxInput `json:"tx,omitempty"`
		Block    *BlockInput      `json:"block,omitempty"`
		CallData hexutil.Bytes    `json:"call_data,omitempty"`
	}
	var dec initInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Tx != nil {
		x.Tx = dec.Tx
	}
	if dec.Block != nil {
		x.Block = dec.Block
	}
	x.CallData = dec.CallData
	return nil
}

// MarshalJSON marshals as JSON.
func (x *PostTxExecuteInput) MarshalJSON() ([]byte, error) {
	type postTxExecuteInput struct {
		Tx      *WithFromTxInput `json:"tx,omitempty"`
		Block   *BlockInput      `json:"block,omitempty"`
		Receipt *ReceiptInput    `json:"receipt,omitempty"`
	}
	var enc postTxExecuteInput
	if x.Tx != nil {
		enc.Tx = x.Tx
	}
	if x.Block != nil {
		enc.Block = x.Block
	}
	if x.Receipt != nil {
		enc.Receipt = x.Receipt
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *PostTxExecuteInput) UnmarshalJSON(input []byte) error {
	type postTxExecuteInput struct {
		Tx      *WithFromTxInput `json:"tx,omitempty"`
		Block   *BlockInput      `json:"block,omitempty"`
		Receipt *ReceiptInput    `json:"receipt,omitempty"`
	}
	var dec postTxExecuteInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Tx != nil {
		x.Tx = dec.Tx
	}
	if dec.Block != nil {
		x.Block = dec.Block
	}
	if dec.Receipt != nil {
		x.Receipt = dec.Receipt
	}
	return nil
}

// MarshalJSON marshals as JSON.
func (x *PreContractCallInput) MarshalJSON() ([]byte, error) {
	type preContractCallInput struct {
		Call  *PreExecMessageInput `json:"call,omitempty"`
		Block *BlockInput          `json:"block,omitempty"`
	}
	var enc preContractCallInput
	if x.Call != nil {
		enc.Call = x.Call
	}
	if x.Block != nil {
		enc.Block = x.Block
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *PreContractCallInput) UnmarshalJSON(input []byte) error {
	type preContractCallInput struct {
		Call  *PreExecMessageInput `json:"call,omitempty"`
		Block *BlockInput          `json:"block,omitempty"`
	}
	var dec preContractCallInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Call != nil {
		x.Call = dec.Call
	}
	if dec.Block != nil {
		x.Block = dec.Block
	}
	return nil
}

// MarshalJSON marshals as JSON.
func (x *PostContractCallInput) MarshalJSON() ([]byte, error) {
	type postContractCallInput struct {
		Call  *PostExecMessageInput `json:"call,omitempty"`
		Block *BlockInput           `json:"block,omitempty"`
	}
	var enc postContractCallInput
	if x.Call != nil {
		enc.Call = x.Call
	}
	if x.Block != nil {
		enc.Block = x.Block
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *PostContractCallInput) UnmarshalJSON(input []byte) error {
	type postContractCallInput struct {
		Call  *PostExecMessageInput `json:"call,omitempty"`
		Block *BlockInput           `json:"block,omitempty"`
	}
	var dec postContractCallInput
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Call != nil {
		x.Call = dec.Call
	}
	if dec.Block != nil {
		x.Block = dec.Block
	}
	return nil
}
