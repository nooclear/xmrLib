package xmrLib

import (
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
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
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_account_tags:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_account_tags",
			Params:  map[string]interface{}{},
		}); err != nil {
		aLog.Error("xmrLib:get_account_tags", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_account_tags:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_account_tags:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
