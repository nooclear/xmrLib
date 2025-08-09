package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type QueryKeyParams struct {
	KeyType keyType `json:"key_type"`
}

type keyType string

var mnemonic keyType = "mnemonic"
var viewkey keyType = "view_key"
var spendkey keyType = "spend_key"

type QueryKeyResult struct {
	Key string `json:"key"`
}

func (wallet *Wallet) QueryKey(id string, params QueryKeyParams) (result QueryKeyResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "query_key",
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
