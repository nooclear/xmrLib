package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type GetReserveProofParams struct {
	All          bool   `json:"all"`
	AccountIndex uint64 `json:"account_index"`
	Amount       uint64 `json:"amount"`
	Message      string `json:"message"`
}

type GetReserveProofResult struct {
	Signature string `json:"signature"`
}

func (wallet *Wallet) GetReserveProof(id string, params GetReserveProofParams) (result GetReserveProofResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_reserve_proof",
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
