package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
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
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:parse_uri:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "parse_uri",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:parse_uri", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:parse_uri:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:parse_uri:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
