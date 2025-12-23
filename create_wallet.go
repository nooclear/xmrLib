package xmrLib

import (
	"encoding/json"
	"errors"

	"github.com/nooclear/jrpcLib"
)

type CreateWalletParams struct {
	Filename string `json:"filename"`
	Password string `json:"password"`
	Language string `json:"language"`
}

func (wallet *Wallet) CreateWallet(id string, params CreateWalletParams) (err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "create_wallet",
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
				return errors.New("create wallet failed")
			}
		}
	}
}
