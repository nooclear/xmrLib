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

func (wallet *Wallet) GetAddressIndex(id string, params AddressIndexParams) (result AddressIndexResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_address_index",
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
