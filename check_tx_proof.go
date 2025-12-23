package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type CheckTxProofParams struct {
	TxID      string `json:"txid"`
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type CheckTxProofResult struct {
	Confirmations uint64 `json:"confirmations"`
	Good          bool   `json:"good"`
	InPool        bool   `json:"in_pool"`
	Received      uint64 `json:"received"`
}

func (wallet *Wallet) CheckTxProof(id string, params CheckTxProofParams) (result CheckTxProofResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_tx_proof",
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
