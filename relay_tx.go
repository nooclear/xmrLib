package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type RelayTxParams struct {
	Hex string `json:"hex"`
}

type RelayTxResult struct {
	TxHash string `json:"tx_hash"`
}

func (wallet *Wallet) RelayTx(id string, params RelayTxParams) (result RelayTxResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "relay_tx",
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
