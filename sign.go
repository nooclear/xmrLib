package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type SignParams struct {
	Data string `json:"data"`
}
type SignResult struct {
	Signature string `json:"signature"`
}

func (wallet *Wallet) Sign(id string, params SignParams) (result SignResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "sign",
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
