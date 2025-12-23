package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type SignMultisigParams struct {
	TxDataHex string `json:"tx_data_hex"`
}

type SignMultisigResult struct {
	TxDataHex  string   `json:"tx_data_hex"`
	TxHashList []string `json:"tx_hash_list"`
}

func (wallet *Wallet) SignMultisig(id string, params SignMultisigParams) (result SignMultisigResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "sign_multisig",
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
