package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type EstimateTxSizeAndWeightParams struct {
	NInputs  uint64 `json:"n_inputs"`
	NOutputs uint64 `json:"n_outputs"`
	RingSize uint64 `json:"ring_size"`
	RCT      bool   `json:"rct"`
}

type EstimateTxSizeAndWeightResult struct {
	Size   uint64 `json:"size"`
	Weight uint64 `json:"weight"`
}

func (wallet *Wallet) EstimateTxSizeAndWeight(id string, params EstimateTxSizeAndWeightParams) (result EstimateTxSizeAndWeightResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:estimate_tx_size_and_weight:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "estimate_tx_size_and_weight",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:estimate_tx_size_and_weight", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:estimate_tx_size_and_weight:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:estimate_tx_size_and_weight:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
