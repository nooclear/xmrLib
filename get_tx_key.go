package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type GetTxKeyParams struct {
	TxID string `json:"txid"`
}

type GetTxKeyResult struct {
	TxKey string `json:"tx_key"`
}

func (wallet *Wallet) GetTxKey(id string, params GetTxKeyParams) (result GetTxKeyResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_tx_key",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return result, mapToStruct(jrpcRes.Result, &result)
		}
	}
}
