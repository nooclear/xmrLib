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

func (wallet *Wallet) CreateAccount(id string, params CreateAccountParams) (result CreateAccountResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "create_account",
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
