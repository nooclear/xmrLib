package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type FinalizeMultisigParams struct {
	MultisigInfo []string `json:"multisig_info"`
	Password     string   `json:"password"`
}

type FinalizeMultisigResult struct {
	Address string `json:"address"`
}

func (wallet *Wallet) FinalizeMultisig(id string, params FinalizeMultisigParams) (result FinalizeMultisigResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:finalize_multisig:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "finalize_multisig",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:finalize_multisig", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:finalize_multisig:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:finalize_multisig:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
