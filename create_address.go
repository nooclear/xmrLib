package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type CreateAddressParams struct {
	AccountIndex uint64 `json:"account_index"`
	Label        string `json:"label"`
	Count        uint64 `json:"count"`
}

type CreateAddressResult struct {
	Address        string   `json:"address"`
	AddressIndex   uint64   `json:"address_index"`
	AddressIndeces []uint64 `json:"address_indices"`
	Addresses      []string `json:"addresses"`
}

func (wallet *Wallet) CreateAddress(id string, params CreateAddressParams) (createaddrResult CreateAddressResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "create_address",
			Params:  convertToMap(json.Marshal(params)),
		}); err != nil {
		return createaddrResult, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return createaddrResult, err
		} else {
			return convertToCreateAddressResult(jrpcRes.Result)
		}
	}
}

func convertToCreateAddressResult(data map[string]interface{}) (result CreateAddressResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
