package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type RefreshParams struct {
	StartHeight uint64 `json:"start_height"`
}
type RefreshResult struct {
	BlocksFetched uint64 `json:"blocks_fetched"`
	ReceivedMoney bool   `json:"received_money"`
}

func (wallet *Wallet) Refresh(id string, params RefreshParams) (result RefreshResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "refresh",
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
