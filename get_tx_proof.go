package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type GetTxProofParams struct {
	TxID    string `json:"txid"`
	Address string `json:"address"`
	Message string `json:"message"`
}

type GetTxProofResult struct {
	Signature string `json:"signature"`
}

func (wallet *Wallet) GetTxProof(id string, params GetTxProofParams) (result GetTxProofResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_tx_proof",
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
