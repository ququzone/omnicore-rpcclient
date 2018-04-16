package rpcclient

import (
	"github.com/btcsuite/btcd/btcjson"
)

// GetTransactionCmd omni_gettransaction
type GetTransactionCmd struct {
	Txid string
}

// NewGetTransactionCmd ...
func NewGetTransactionCmd(txHash string) *GetTransactionCmd {
	return &GetTransactionCmd{
		Txid: txHash,
	}
}

// ListBlockTransactionsCmd omni_listblocktransactions
type ListBlockTransactionsCmd struct {
	Index int64
}

// NewListBlockTransactionsCmd ...
func NewListBlockTransactionsCmd(index int64) *ListBlockTransactionsCmd {
	return &ListBlockTransactionsCmd{
		Index: index,
	}
}

// GetBalanceCmd omni_getbalance
type GetBalanceCmd struct {
	Address    string
	PropertyID int64
}

// NewGetBalanceCmd ...
func NewGetBalanceCmd(address string, propertyID int64) *GetBalanceCmd {
	return &GetBalanceCmd{
		Address:    address,
		PropertyID: propertyID,
	}
}

func init() {
	flags := btcjson.UFWalletOnly
	btcjson.MustRegisterCmd("omni_gettransaction", (*GetTransactionCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_listblocktransactions", (*ListBlockTransactionsCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_getbalance", (*GetBalanceCmd)(nil), flags)
}
