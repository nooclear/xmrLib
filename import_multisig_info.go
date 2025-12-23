package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type ImportMultisigInfoParams struct {
	Info []string `json:"info"`
}

type ImportMultisigInfoResult struct {
	NOutputs uint64 `json:"n_outputs"`
}

func (wallet *Wallet) ImportMultisigInfo(id string, params ImportMultisigInfoParams) (result ImportMultisigInfoResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "import_multisig_info",
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
