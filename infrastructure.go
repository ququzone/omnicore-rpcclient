package rpcclient

import (
	"encoding/json"
)

// jsonRequest holds information about a json request that is used to properly
// detect, interpret, and deliver a reply to it.
type jsonRequest struct {
	id             uint64
	method         string
	cmd            interface{}
	marshalledJSON []byte
}

// Response ...
type Response struct {
	ID     int32           `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Transaction ...
type Transaction struct {
	TXID             string `json:"txid"`
	Fee              string `json:"fee"`
	SendingAddress   string `json:"sendingaddress"`
	ReferenceAddress string `json:"referenceaddress"`
	IsMine           bool   `json:"ismine"`
	Version          int32  `json:"version"`
	TypeInt          int32  `json:"type_int"`
	Type             string `json:"type"`
	PropertyID       int64  `json:"propertyid"`
	Divisible        bool   `json:"divisible"`
	Amount           string `json:"amount"`
	Valid            bool   `json:"valid"`
	BlockHash        string `json:"blockhash"`
	BlockTime        int64  `json:"blocktime"`
	PositionInBlock  int32  `json:"positioninblock"`
	Block            int64  `json:"block"`
	Confirmations    int64  `json:"confirmations"`
}

// Balance ...
type Balance struct {
	Balance  string `json:"balance,omitempty"`
	Reserved string `json:"reserved,omitempty"`
	Frozen   string `json:"frozen,omitempty"`
}
