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
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return convertToQueryKeyResult(jrpcRes.Result)
		}
	}
}

func convertToQueryKeyResult(data map[string]interface{}) (result QueryKeyResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
