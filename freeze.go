package xmrLib

import (
	"encoding/json"
	"errors"

	"github.com/nooclear/jrpcLib"
)

type FreezeParams struct {
	KeyImage string `json:"key_image"`
}

func (wallet *Wallet) Freeze(id string, params FreezeParams) (err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "freeze",
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
				return errors.New("freeze failed")
			}
		}
	}
}
