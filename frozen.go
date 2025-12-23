package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type FrozenParams struct {
	KeyImage string `json:"key_image"`
}

type FrozenResult struct {
	Frozen bool `json:"frozen"`
}

func (wallet *Wallet) Frozen(id string, params FrozenParams) (result FrozenResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "frozen",
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
