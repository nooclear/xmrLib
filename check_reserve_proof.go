package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type CheckReserveProofParams struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type CheckReserveProofResult struct {
	Good  bool   `json:"good"`
	Spent uint64 `json:"spent"`
	Total uint64 `json:"total"`
}

func (wallet *Wallet) CheckReserveProof(id string, params CheckReserveProofParams) (result CheckReserveProofResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_reserve_proof",
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
