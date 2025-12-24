package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type GetTransfersParams struct {
	In             bool     `json:"in"`
	Out            bool     `json:"out"`
	Pending        bool     `json:"pending"`
	Failed         bool     `json:"failed"`
	Pool           bool     `json:"pool"`
	FilterByHeight bool     `json:"filter_by_height"`
	MinHeight      uint64   `json:"min_height"`
	MaxHeight      uint64   `json:"max_height"`
	AccountIndex   uint64   `json:"account_index"`
	SubaddrIndices []uint64 `json:"subaddr_indices"`
	AllAccounts    bool     `json:"all_accounts"`
}

type GetTransfersResult struct {
	In []struct {
		Address         string   `json:"address"`
		Amount          uint64   `json:"amount"`
		Amounts         []uint64 `json:"amounts"`
		Confirmations   uint64   `json:"confirmations"`
		DoubleSpendSeen bool     `json:"double_spend_seen"`
		Fee             uint64   `json:"fee"`
		Height          uint64   `json:"height"`
		Note            string   `json:"note"`
		Destinations    []struct {
			Amount  uint64 `json:"amount"`
			Address string `json:"address"`
		} `json:"destinations"`
		PaymentID    string `json:"payment_id"`
		SubaddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		SubAddrIndices []struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_indices"`
		SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
		Timestamp                       uint64 `json:"timestamp"`
		TxID                            string `json:"txid"`
		Type                            string `json:"type"`
		UnlockTime                      uint64 `json:"unlock_time"`
		Locked                          bool   `json:"locked"`
	} `json:"in"`
	Out []struct {
		Address         string   `json:"address"`
		Amount          uint64   `json:"amount"`
		Amounts         []uint64 `json:"amounts"`
		Confirmations   uint64   `json:"confirmations"`
		DoubleSpendSeen bool     `json:"double_spend_seen"`
		Fee             uint64   `json:"fee"`
		Height          uint64   `json:"height"`
		Note            string   `json:"note"`
		Destinations    []struct {
			Amount  uint64 `json:"amount"`
			Address string `json:"address"`
		} `json:"destinations"`
		PaymentID    string `json:"payment_id"`
		SubaddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		SubAddrIndices []struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_indices"`
		SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
		Timestamp                       uint64 `json:"timestamp"`
		TxID                            string `json:"txid"`
		Type                            string `json:"type"`
		UnlockTime                      uint64 `json:"unlock_time"`
		Locked                          bool   `json:"locked"`
	} `json:"out"`
	Pending []struct {
		Address         string   `json:"address"`
		Amount          uint64   `json:"amount"`
		Amounts         []uint64 `json:"amounts"`
		Confirmations   uint64   `json:"confirmations"`
		DoubleSpendSeen bool     `json:"double_spend_seen"`
		Fee             uint64   `json:"fee"`
		Height          uint64   `json:"height"`
		Note            string   `json:"note"`
		Destinations    []struct {
			Amount  uint64 `json:"amount"`
			Address string `json:"address"`
		} `json:"destinations"`
		PaymentID    string `json:"payment_id"`
		SubaddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		SubAddrIndices []struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_indices"`
		SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
		Timestamp                       uint64 `json:"timestamp"`
		TxID                            string `json:"txid"`
		Type                            string `json:"type"`
		UnlockTime                      uint64 `json:"unlock_time"`
		Locked                          bool   `json:"locked"`
	} `json:"pending"`
	Failed []struct {
		Address         string   `json:"address"`
		Amount          uint64   `json:"amount"`
		Amounts         []uint64 `json:"amounts"`
		Confirmations   uint64   `json:"confirmations"`
		DoubleSpendSeen bool     `json:"double_spend_seen"`
		Fee             uint64   `json:"fee"`
		Height          uint64   `json:"height"`
		Note            string   `json:"note"`
		Destinations    []struct {
			Amount  uint64 `json:"amount"`
			Address string `json:"address"`
		} `json:"destinations"`
		PaymentID    string `json:"payment_id"`
		SubaddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		SubAddrIndices []struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_indices"`
		SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
		Timestamp                       uint64 `json:"timestamp"`
		TxID                            string `json:"txid"`
		Type                            string `json:"type"`
		UnlockTime                      uint64 `json:"unlock_time"`
		Locked                          bool   `json:"locked"`
	} `json:"failed"`
	Pool []struct {
		Address         string   `json:"address"`
		Amount          uint64   `json:"amount"`
		Amounts         []uint64 `json:"amounts"`
		Confirmations   uint64   `json:"confirmations"`
		DoubleSpendSeen bool     `json:"double_spend_seen"`
		Fee             uint64   `json:"fee"`
		Height          uint64   `json:"height"`
		Note            string   `json:"note"`
		Destinations    []struct {
			Amount  uint64 `json:"amount"`
			Address string `json:"address"`
		} `json:"destinations"`
		PaymentID    string `json:"payment_id"`
		SubaddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		SubAddrIndices []struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_indices"`
		SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
		Timestamp                       uint64 `json:"timestamp"`
		TxID                            string `json:"txid"`
		Type                            string `json:"type"`
		UnlockTime                      uint64 `json:"unlock_time"`
		Locked                          bool   `json:"locked"`
	} `json:"pool"`
}

func (wallet *Wallet) GetTransfers(id string, params GetTransfersParams) (result GetTransfersResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_transfers:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "get_transfers",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:get_transfers", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_transfers:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_transfers:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
