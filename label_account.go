package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

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
			Params:  convertToMap(json.Marshal(params)),
		}); err != nil {
		return err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return err
		} else {
			if len(jrpcRes.Result) == 0 {
				return nil
			} else {
				fmt.Println(jrpcRes.Result)
				return errors.New("label account failed")
			}
		}
	}
}
