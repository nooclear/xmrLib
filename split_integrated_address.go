package xmrLib

import (
	"encoding/json"

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
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "split_integrated_address",
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
