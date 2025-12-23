package xmrLib

import (
	"encoding/json"

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
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "estimate_tx_size_and_weight",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return result, mapToStruct(jrpcRes.Result, &result)
		}
	}
}
