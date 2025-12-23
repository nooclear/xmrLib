package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type CheckSpendProofParams struct {
	TxID      string `json:"txid"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type CheckSpendProofResult struct {
	Good bool `json:"good"`
}

func (wallet *Wallet) CheckSpendProof(id string, params CheckSpendProofParams) (result CheckSpendProofResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_spend_proof",
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
