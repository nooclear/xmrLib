package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type GetSpendProofParams struct {
	TxID    string `json:"txid"`
	Message string `json:"message"`
}

type GetSpendProofResult struct {
	Signature string `json:"signature"`
}

func (wallet *Wallet) GetSpendProof(id string, params GetSpendProofParams) (result GetSpendProofResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_spend_proof",
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
