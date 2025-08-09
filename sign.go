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
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return convertToSignResult(jrpcRes.Result)
		}
	}
}

func convertToSignResult(data map[string]interface{}) (result SignResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
