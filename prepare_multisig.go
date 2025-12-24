package xmrLib

import (
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type PrepareMultisigResult struct {
	MultisigInfo string `json:"multisig_info"`
}

func (wallet *Wallet) PrepareMultisig(id string) (result PrepareMultisigResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:prepare_multisig:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "prepare_multisig",
		Params:  map[string]interface{}{},
	}); err != nil {
		aLog.Error("xmrLib:prepare_multisig", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:prepare_multisig:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:prepare_multisig:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
