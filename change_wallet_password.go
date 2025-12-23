package xmrLib

import (
	"encoding/json"
	"errors"

	"github.com/nooclear/jrpcLib"
)

type ChangeWalletPasswordParams struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (wallet *Wallet) ChangeWalletPassword(id string, params ChangeWalletPasswordParams) (err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "change_wallet_password",
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
				return errors.New("change wallet password failed")
			}
		}
	}
}
