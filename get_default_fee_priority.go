package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type DefaultFeePriorityResult struct {
	FeePriority uint64 `json:"fee_priority"`
}

func (wallet *Wallet) GetDefaultFeePriority(id string) (defRes DefaultFeePriorityResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_default_fee_priority",
			Params:  nil,
		}); err != nil {
		return defRes, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return defRes, err
		} else {
			return convertToDefaultFeePriorityResult(jrpcRes.Result)
		}
	}
}

func convertToDefaultFeePriorityResult(data map[string]interface{}) (result DefaultFeePriorityResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
