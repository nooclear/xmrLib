package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type CreateAccountParams struct {
	Label string `json:"label"`
}

type CreateAccountResult struct {
	AccountIndex uint64 `json:"account_index"`
	Address      string `json:"address"`
}

func (wallet *Wallet) CreateAccount(id string, params CreateAccountParams) (createResult CreateAccountResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "create_account",
			Params:  convertToMap(json.Marshal(params)),
		}); err != nil {
		return createResult, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return createResult, err
		} else {
			return convertToCreateAccountResult(jrpcRes.Result)
		}
	}
}

func convertToCreateAccountResult(data map[string]interface{}) (result CreateAccountResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
