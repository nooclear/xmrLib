package xmrLib

import (
	"encoding/json"
	"errors"

	"github.com/nooclear/jrpcLib"
)

type TagAccountsParams struct {
	Tag      string   `json:"tags"`
	Accounts []uint64 `json:"accounts"`
}

func (wallet *Wallet) TagAccounts(id string, params TagAccountsParams) (err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "tag_accounts",
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
				return errors.New("tag accounts failed")
			}
		}
	}
}
