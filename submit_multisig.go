package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type SubmitMultisigParams struct {
	TxDataHex string `json:"tx_data_hex"`
}

type SubmitMultisigResult struct {
	TxHashList []string `json:"tx_hash_list"`
}

func (wallet *Wallet) SubmitMultisig(id string, params SubmitMultisigParams) (result SubmitMultisigResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:submit_multisig:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "submit_multisig",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:submit_multisig", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:submit_multisig:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:submit_multisig:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
