package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type GetTransferByTxIDParams struct {
	TxID         string `json:"txid"`
	AccountIndex uint64 `json:"account_index"`
}

type GetTransferByTxIDResult struct {
	Transfers []struct {
		Transfer struct {
			Address       string `json:"address"`
			Amount        uint64 `json:"amount"`
			Amounts       any    `json:"amounts"`
			Confirmations uint64 `json:"confirmations"`
			Destinations  struct {
				Amount  uint64 `json:"amount"`
				Address string `json:"address"`
			} `json:"destinations"`
			DoubleSpendSeen bool   `json:"double_spend_seen"`
			Fee             uint64 `json:"fee"`
			Height          uint64 `json:"height"`
			Locked          bool   `json:"locked"`
			Note            string `json:"note"`
			PaymentID       string `json:"payment_id"`
			SubaddrIndex    struct {
				Major uint64 `json:"major"`
				Minor uint64 `json:"minor"`
			} `json:"subaddr_index"`
			SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
			Timestamp                       uint64 `json:"timestamp"`
			TxID                            string `json:"txid"`
			Type                            string `json:"type"`
			UnlockTime                      uint64 `json:"unlock_time"`
		} `json:"transfer"`
	} `json:"transfers"`
}

func (wallet *Wallet) GetTransferByTxID(id string, params GetTransferByTxIDParams) (result GetTransferByTxIDResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_transfer_by_tx_id:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_transfer_by_tx_id",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:get_transfer_by_tx_id", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_transfer_by_tx_id:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_transfer_by_tx_id:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
