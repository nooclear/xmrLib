package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type SubmitMultisigParams struct {
	TxDataHex string `json:"tx_data_hex"`
}

type SubmitMultisigResult struct {
	TxHashList []string `json:"tx_hash_list"`
}

func (wallet *Wallet) SubmitMultisig(id string, params SubmitMultisigParams) (result SubmitMultisigResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "submit_multisig",
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
