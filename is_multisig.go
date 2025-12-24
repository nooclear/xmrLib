package xmrLib

import (
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type IsMultisigResult struct {
	Multisig  bool   `json:"multisig"`
	Ready     bool   `json:"ready"`
	Threshold uint64 `json:"threshold"`
	Total     uint64 `json:"total"`
}

func (wallet *Wallet) IsMultisig(id string) (result IsMultisigResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:is_multisig:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "is_multisig",
		Params:  map[string]interface{}{},
	}); err != nil {
		aLog.Error("xmrLib:is_multisig", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:is_multisig:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:is_multisig:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
