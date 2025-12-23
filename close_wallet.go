package xmrLib

import (
	"errors"

	"github.com/nooclear/jrpcLib"
)

func (wallet *Wallet) CloseWallet(id string) (err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "close_wallet",
		Params:  map[string]interface{}{},
	}); err != nil {
		return err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return err
		} else {
			if len(jrpcRes.Result) == 0 {
				return nil
			} else {
				return errors.New("close wallet failed")
			}
		}
	}
}
