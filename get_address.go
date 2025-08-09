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

func (wallet *Wallet) GetAddress(id string, params AddressParams) (result AddressResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_address",
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
