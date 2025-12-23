package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type SignTransferParams struct {
}

type SignTransferResult struct {
}

func (wallet *Wallet) SignTransfer(id string, params SignTransferParams) (result SignTransferResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "sign_transfer",
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
