package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type GetBulkPaymentsParams struct {
	PaymentIDs     []string `json:"payment_ids"`
	MinBlockHeight uint64   `json:"min_block_height"`
}

type GetBulkPaymentsResult struct {
	Payments []struct {
		PaymentID    string `json:"payment_id"`
		TxHash       string `json:"tx_hash"`
		Amount       uint64 `json:"amount"`
		BlockHeight  uint64 `json:"block_height"`
		UnlockTime   uint64 `json:"unlock_time"`
		SubaddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		Address string `json:"address"`
	}
}

func (wallet *Wallet) GetBulkPayments(id string, params GetBulkPaymentsParams) (result GetBulkPaymentsResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_bulk_payments:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_bulk_payments",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:get_bulk_payments", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_bulk_payments:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_bulk_payments:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
