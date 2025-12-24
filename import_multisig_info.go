package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ImportMultisigInfoParams struct {
	Info []string `json:"info"`
}

type ImportMultisigInfoResult struct {
	NOutputs uint64 `json:"n_outputs"`
}

func (wallet *Wallet) ImportMultisigInfo(id string, params ImportMultisigInfoParams) (result ImportMultisigInfoResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:import_multisig_info:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "import_multisig_info",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:import_multisig_info", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:import_multisig_info:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:import_multisig_info:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
