package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nooclear/jrpcLib"
)

type LabelAddressParams struct {
	Index struct {
		Major uint64 `json:"major"`
		Minor uint64 `json:"minor"`
	} `json:"index"`
	Label string `json:"label"`
}

func (wallet *Wallet) LabelAddress(id string, params LabelAddressParams) (err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "label_address",
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
				return errors.New("label address failed")
			}
		}
	}
}
