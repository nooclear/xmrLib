package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type FinalizeMultisigParams struct {
	MultisigInfo []string `json:"multisig_info"`
	Password     string   `json:"password"`
}

type FinalizeMultisigResult struct {
	Address string `json:"address"`
}

func (wallet *Wallet) FinalizeMultisig(id string, params FinalizeMultisigParams) (result FinalizeMultisigResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "finalize_multisig",
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
