package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

type ExportMultisigInfoResult struct {
	Info string `json:"info"`
}

func (wallet *Wallet) ExportMultisigInfo(id string) (result ExportMultisigInfoResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "export_multisig_info",
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
