package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type MakeMultisigParams struct {
	MultisigInfo []string `json:"multisig_info"`
	Threshold    uint64   `json:"threshold"`
	Password     string   `json:"password"`
}

type MakeMultisigResult struct {
	Address      string `json:"address"`
	MultisigInfo string `json:"multisig_info"`
}

func (wallet *Wallet) MakeMultisig(id string, params MakeMultisigParams) (result MakeMultisigResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "make_multisig",
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
