package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type CheckSpendProofParams struct {
	TxID      string `json:"txid"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type CheckSpendProofResult struct {
	Good bool `json:"good"`
}

func (wallet *Wallet) CheckSpendProof(id string, params CheckSpendProofParams) (result CheckSpendProofResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:check_spend_proof:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_spend_proof",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:check_spend_proof", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:check_spend_proof:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:check_spend_proof:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
