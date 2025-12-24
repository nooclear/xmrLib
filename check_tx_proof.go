package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type CheckTxProofParams struct {
	TxID      string `json:"txid"`
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type CheckTxProofResult struct {
	Confirmations uint64 `json:"confirmations"`
	Good          bool   `json:"good"`
	InPool        bool   `json:"in_pool"`
	Received      uint64 `json:"received"`
}

func (wallet *Wallet) CheckTxProof(id string, params CheckTxProofParams) (result CheckTxProofResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:check_tx_proof:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_tx_proof",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:check_tx_proof", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:check_tx_proof:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:check_tx_proof:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}

}
