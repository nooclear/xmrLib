package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
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
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:sign_multisig:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "sign_multisig",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:sign_multisig", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:sign_multisig:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:sign_multisig:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
