package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type GetAttributeParams struct {
	Key string `json:"key"`
}
type GetAttributeResult struct {
	Value string `json:"value"`
}

func (wallet *Wallet) GetAttribute(id string, params GetAttributeParams) (result GetAttributeResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_attribute",
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
