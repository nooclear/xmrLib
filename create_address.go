package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
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

func (wallet *Wallet) CreateAddress(id string, params CreateAddressParams) (result CreateAddressResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:create_address:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "create_address",
			Params:  bytesToMap(json.Marshal(params)),
		}); err != nil {
		aLog.Error("xmrLib:create_address", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:create_address:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:create_address:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
