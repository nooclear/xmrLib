package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

type GetAccountTagsResponse struct {
	AccountTags []struct {
		Tag      string   `json:"tag"`
		Label    string   `json:"label"`
		Accounts []uint64 `json:"accounts"`
	} `json:"account_tags"`
}

func (wallet *Wallet) GetAccountTags(id string) (result GetAccountTagsResponse, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_account_tags",
			Params:  map[string]interface{}{},
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
