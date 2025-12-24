package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
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
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:exchange_multisig_keys:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "exchange_multisig_keys",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:exchange_multisig_keys", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:exchange_multisig_keys:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:exchange_multisig_keys:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
