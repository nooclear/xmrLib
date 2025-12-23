package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type GetTxNotesParams struct {
	TxIDs []string `json:"tx_ids"`
}

type GetTxNotesResult struct {
	Notes []string `json:"notes"`
}

func (wallet *Wallet) GetTxNotes(id string, params GetTxNotesParams) (result GetTxNotesResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_tx_notes",
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
