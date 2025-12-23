package xmrLib

import (
	"encoding/json"
	"errors"

	"github.com/nooclear/jrpcLib"
)

type SetTxNotesParams struct {
	TxIDs []string `json:"tx_ids"`
	Notes []string `json:"notes"`
}

func (wallet *Wallet) SetTxNotes(id string, params SetTxNotesParams) (err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "set_tx_notes",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		return err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return err
		} else {
			if len(jrpcRes.Result) == 0 {
				return nil
			} else {
				return errors.New("set tx notes failed")
			}
		}
	}
}
