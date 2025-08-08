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

func (wallet *Wallet) SplitIntegratedAddress(id string, params SplitIntegratedAddressParams) (siaRes SplitIntegratedAddressResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "split_integrated_address",
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return siaRes, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return siaRes, err
		} else {
			return convertToSplitIntegratedAddressResult(jrpcRes.Result)
		}
	}
}

func convertToSplitIntegratedAddressResult(data map[string]interface{}) (result SplitIntegratedAddressResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
