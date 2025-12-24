package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
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

func (wallet *Wallet) MakeUri(id string, params MakeUriParams) (result MakeUriResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:make_uri:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "make_uri",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:make_uri", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:make_uri:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:make_uri:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
