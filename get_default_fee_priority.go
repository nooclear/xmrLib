package xmrLib

import (
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type DefaultFeePriorityResult struct {
	FeePriority uint64 `json:"fee_priority"`
}

func (wallet *Wallet) GetDefaultFeePriority(id string) (result DefaultFeePriorityResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:getDefaultFeePriority:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_default_fee_priority",
			Params:  nil,
		}); err != nil {
		aLog.Error("xmrLib:getDefaultFeePriority", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:getDefaultFeePriority:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:getDefaultFeePriority:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
