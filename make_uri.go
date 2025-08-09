package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type MakeUriParams struct {
	Address       string `json:"address"`
	Amount        uint64 `json:"amount"`
	PaymentID     string `json:"payment_id"`
	RecipientName string `json:"recipient_name"`
	TxDescription string `json:"tx_description"`
}

type MakeUriResult struct {
	URI string `json:"uri"`
}

func (wallet *Wallet) MakeUri(id string, params MakeUriParams) (uriRes MakeUriResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "make_uri",
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return uriRes, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return uriRes, err
		} else {
			return convertToMakeUriResult(jrpcRes.Result)
		}
	}
}

func convertToMakeUriResult(data map[string]interface{}) (result MakeUriResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
