package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type MakeIntegratedAddressParams struct {
	PaymentID       string `json:"payment_id"`
	StandardAddress string `json:"standard_address"`
}

type MakeIntegratedAddressResult struct {
	IntegratedAddress string `json:"integrated_address"`
	PaymentID         string `json:"payment_id"`
}

func (wallet *Wallet) MakeIntegratedAddress(id string, params MakeIntegratedAddressParams) (result MakeIntegratedAddressResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:make_integrated_address:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "make_integrated_address",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:make_integrated_address", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:make_integrated_address:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:make_integrated_address:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
