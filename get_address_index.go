package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
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
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_address_index:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_address_index",
			Params:  bytesToMap(json.Marshal(params)),
		}); err != nil {
		aLog.Error("xmrLib:get_address_index", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_address_index:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_address_index:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
