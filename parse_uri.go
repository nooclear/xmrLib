package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type ParseUriParams struct {
	URI string `json:"uri"`
}

type ParseUriResult struct {
	URI struct {
		Address       string `json:"address"`
		Amount        uint64 `json:"amount"`
		PaymentID     string `json:"payment_id"`
		RecipientName string `json:"recipient_name"`
		TxDescription string `json:"tx_description"`
	} `json:"uri"`
}

func (wallet *Wallet) ParseURI(id string, params ParseUriParams) (result ParseUriResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "parse_uri",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return result, mapToStruct(jrpcRes.Result, &result)
		}
	}
}
