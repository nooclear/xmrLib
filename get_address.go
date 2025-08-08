package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type AddressParams struct {
	AccountIndex uint64   `json:"account_index"`
	AddressIndex []uint64 `json:"address_index"`
}

type AddressResult struct {
	Address   string `json:"address"`
	Addresses []struct {
		Address      string `json:"address"`
		AddressIndex uint64 `json:"address_index"`
		Label        string `json:"label"`
		Used         bool   `json:"used"`
	} `json:"addresses"`
}

func (wallet *Wallet) GetAddress(id string, params AddressParams) (addrRes AddressResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_address",
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return addrRes, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return addrRes, err
		} else {
			return convertToAddressResult(jrpcRes.Result)
		}
	}
}

func convertToAddressResult(data map[string]interface{}) (result AddressResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
