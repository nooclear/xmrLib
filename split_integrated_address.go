package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type SplitIntegratedAddressParams struct {
	IntegratedAddress string `json:"integrated_address"`
}

type SplitIntegratedAddressResult struct {
	IsSubaddress    bool   `json:"is_subaddress"`
	PaymentID       string `json:"payment_id"`
	StandardAddress string `json:"standard_address"`
}

func (wallet *Wallet) SplitIntegratedAddress(id string, params SplitIntegratedAddressParams) (result SplitIntegratedAddressResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:split_integrated_address:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "split_integrated_address",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:split_integrated_address", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:split_integrated_address:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:split_integrated_address:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
