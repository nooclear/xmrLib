package xmrLib

import (
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ExportMultisigInfoResult struct {
	Info string `json:"info"`
}

func (wallet *Wallet) ExportMultisigInfo(id string) (result ExportMultisigInfoResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:export_multisig_info:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "export_multisig_info",
		Params:  map[string]interface{}{},
	}); err != nil {
		aLog.Error("xmrLib:export_multisig_info", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:export_multisig_info:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:export_multisig_info:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
