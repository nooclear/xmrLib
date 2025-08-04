package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

func (wallet *Wallet) GetDefaultFeePriority(id string) ([]byte, error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_default_fee_priority",
			Params:  nil,
		}); err != nil {
		return nil, err
	} else {
		defer func() {
			err = res.Body.Close() // need to find a way to properly handle these errors
		}()
		return checkStatus(res)
	}
}
