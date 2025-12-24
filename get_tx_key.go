package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type GetTxKeyParams struct {
	TxID string `json:"txid"`
}

type GetTxKeyResult struct {
	TxKey string `json:"tx_key"`
}

func (wallet *Wallet) GetTxKey(id string, params GetTxKeyParams) (result GetTxKeyResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_tx_key:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_tx_key",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:get_tx_key", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_tx_key:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_tx_key:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
