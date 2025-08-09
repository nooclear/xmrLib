package xmrLib

import (
	"encoding/json"
	"fmt"

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
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			fmt.Println(jrpcRes.Result)
			return convertToAccountTagsResult(jrpcRes.Result)
		}
	}
}

func convertToAccountTagsResult(data map[string]interface{}) (result GetAccountTagsResponse, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		fmt.Println(string(bytes))
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
