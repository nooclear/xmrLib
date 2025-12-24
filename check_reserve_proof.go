package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type CheckReserveProofParams struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type CheckReserveProofResult struct {
	Good  bool   `json:"good"`
	Spent uint64 `json:"spent"`
	Total uint64 `json:"total"`
}

func (wallet *Wallet) CheckReserveProof(id string, params CheckReserveProofParams) (result CheckReserveProofResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:check_reserve_proof:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "check_reserve_proof",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:check_reserve_proof", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:check_reserve_proof:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:check_reserve_proof:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
