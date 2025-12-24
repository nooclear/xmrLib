package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ValidateAddressParams struct {
	Address        string `json:"address"`
	AnyNetType     bool   `json:"any_net_type"`
	AllowOpenAlias bool   `json:"allow_openalias"`
}

type ValidateAddressResult struct {
	Valid            bool   `json:"valid"`
	Integrated       bool   `json:"integrated"`
	Subaddress       bool   `json:"subaddress"`
	NetType          string `json:"net_type"`
	OpenAliasAddress string `json:"openalias_address"`
}

func (wallet *Wallet) ValidateAddress(id string, params ValidateAddressParams) (result ValidateAddressResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:validate_address:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "validate_address",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:validate_address", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:validate_address:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:validate_address:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
