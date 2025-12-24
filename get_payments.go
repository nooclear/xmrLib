package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type GetPaymentsParams struct {
	PaymentID string `json:"payment_id"`
}

type GetPaymentsResult struct {
	Payments []struct {
		PaymentID    string `json:"payment_id"`
		TxHash       string `json:"tx_hash"`
		Amount       uint64 `json:"amount"`
		BlockHeight  uint64 `json:"block_height"`
		UnlockTime   uint64 `json:"unlock_time"`
		Locked       bool   `json:"locked"`
		SubaddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		Address string `json:"address"`
	} `json:"payments"`
}

func (wallet *Wallet) GetPayments(id string, params GetPaymentsParams) (result GetPaymentsResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_payments:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_payments",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:get_payments", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_payments:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_payments:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
