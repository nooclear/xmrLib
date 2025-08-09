package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

type DefaultFeePriorityResult struct {
	FeePriority uint64 `json:"fee_priority"`
}

func (wallet *Wallet) GetDefaultFeePriority(id string) (result DefaultFeePriorityResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_default_fee_priority",
			Params:  nil,
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
