package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type ExchangeMultisigKeysParams struct {
	Password                  string `json:"password"`
	MultisigInfo              string `json:"multisig_info"`
	ForceUpdateUseWithCaution bool   `json:"force_update_use_with_caution"`
}

type ExchangeMultisigKeysResult struct {
	Address      string `json:"address"`
	MultisigInfo string `json:"multisig_info"`
}

func (wallet *Wallet) ExchangeMultisigKeys(id string, params ExchangeMultisigKeysParams) (result ExchangeMultisigKeysResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "exchange_multisig_keys",
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
