package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type MakeIntegratedAddressParams struct {
	StandardAddress string `json:"standard_address"`
	PaymentID       string `json:"payment_id"`
}
type MakeIntegratedAddressResult struct {
	IntegratedAddress string `json:"integrated_address"`
	PaymentID         string `json:"payment_id"`
}

func (wallet *Wallet) MakeIntegratedAddress(id string, params MakeIntegratedAddressParams) (miaRes MakeIntegratedAddressResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "make_integrated_address",
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return miaRes, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return miaRes, err
		} else {
			return convertToIntegratedAddressResult(jrpcRes.Result)
		}
	}
}

func convertToIntegratedAddressResult(data map[string]interface{}) (result MakeIntegratedAddressResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
