package xmrLib

import (
	"encoding/json"

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
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "make_integrated_address",
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
