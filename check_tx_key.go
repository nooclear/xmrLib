package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type CheckTxKeyParams struct {
	TxID    string `json:"txid"`
	TxKey   string `json:"tx_key"`
	Address string `json:"address"`
}

type CheckTxKeyResult struct {
	Confirmations uint64 `json:"confirmations"`
	InPool        bool   `json:"in_pool"`
	Received      uint64 `json:"received"`
}

func (wallet *Wallet) CheckTxKey(id string, params CheckTxKeyParams) (result CheckTxKeyResult, err error) {
	if res, err := wallet.Request(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_tx_key",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		return result, mapToStruct(res.Result, &result)
	}
}
