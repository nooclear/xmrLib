package xmrLib

import "github.com/nooclear/jrpcLib"

type IsMultisigResult struct {
	Multisig  bool   `json:"multisig"`
	Ready     bool   `json:"ready"`
	Threshold uint64 `json:"threshold"`
	Total     uint64 `json:"total"`
}

func (wallet *Wallet) IsMultisig(id string) (result IsMultisigResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "is_multisig",
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
