package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type AddressIndexParams struct {
	Address string `json:"address"`
}

type AddressIndexResult struct {
	Index struct {
		Major uint64 `json:"major"`
		Minor uint64 `json:"minor"`
	} `json:"index"`
}

func (wallet *Wallet) GetAddressIndex(id string, params AddressIndexParams) (addrRes AddressIndexResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_address_index",
			Params:  convertToMap(json.Marshal(params)),
		}); err != nil {
		return addrRes, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return addrRes, err
		} else {
			return convertToAddressIndexResult(jrpcRes.Result)
		}
	}
}

func convertToAddressIndexResult(data map[string]interface{}) (result AddressIndexResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
