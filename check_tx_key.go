package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type CheckTxKeyParams struct {
	TxID    string `json:"txid"`
	TxKey   string `json:"tx_key"`
	Address string `json:"address"`
}

type CheckTxKeyResult struct {
	Confirmations uint64 `json:"confirmations"`
	InPool        bool   `json:"in_pool"`
	Received      uint64 `json:"received"`
}

func (wallet *Wallet) CheckTxKey(id string, params CheckTxKeyParams) (result CheckTxKeyResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:check_tx_key:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Request(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_tx_key",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:check_tx_key", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:check_tx_key:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(res.Result, &result)
	}
}
