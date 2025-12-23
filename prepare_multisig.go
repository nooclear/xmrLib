package xmrLib

import "github.com/nooclear/jrpcLib"

type PrepareMultisigResult struct {
	MultisigInfo string `json:"multisig_info"`
}

func (wallet *Wallet) PrepareMultisig(id string) (result PrepareMultisigResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "prepare_multisig",
		Params:  map[string]interface{}{},
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
