package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type CreateAccountParams struct {
	Label string `json:"label"`
}

type CreateAccountResult struct {
	AccountIndex uint64 `json:"account_index"`
	Address      string `json:"address"`
}

func (wallet *Wallet) CreateAccount(id string, params CreateAccountParams) (result CreateAccountResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:create_account:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "create_account",
			Params:  bytesToMap(json.Marshal(params)),
		}); err != nil {
		aLog.Error("xmrLib:create_account", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:create_account:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:create_account:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
