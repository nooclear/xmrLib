package xmrLib

import (
	"encoding/json"
	"errors"

	"github.com/nooclear/jrpcLib"
)

type LabelAccountParams struct {
	AccountIndex uint64 `json:"account_index"`
	Label        string `json:"label"`
}

func (wallet *Wallet) LabelAccount(id string, params LabelAccountParams) (err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "label_account",
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
				return errors.New("label account failed")
			}
		}
	}
}
