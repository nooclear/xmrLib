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
		Params:  convertToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return convertToParseURIResult(jrpcRes.Result)
		}
	}
}

func convertToParseURIResult(data map[string]interface{}) (result ParseUriResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
