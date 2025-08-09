package xmrLib

import (
	"encoding/json"

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
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "validate_address",
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return convertToValidateAddressResult(jrpcRes.Result)
		}
	}
}

func convertToValidateAddressResult(data map[string]interface{}) (result ValidateAddressResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
