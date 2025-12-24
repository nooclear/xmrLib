package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type IncomingTransfersParams struct {
	TransferType   Transfer `json:"transfer_type"`
	AccountIndex   uint64   `json:"account_index"`
	SubaddrIndices []uint64 `json:"subaddr_indices"`
}

type Transfer string

// Todo
var All = Transfer("all")
var Available = Transfer("available")
var Unavailable = Transfer("unavailable")

type IncomingTransfersResult struct {
	Transfers []struct {
		Amount       uint64 `json:"amount"`
		GlobalIndex  uint64 `json:"global_index"`
		KeyImage     string `json:"key_image"`
		Spent        bool   `json:"spent"`
		SubAddrIndex struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
		} `json:"subaddr_index"`
		TxHash      string `json:"tx_hash"`
		Frozen      bool   `json:"frozen"`
		Unlocked    bool   `json:"unlocked"`
		BlockHeight uint64 `json:"block_height"`
		PubKey      string `json:"pub_key"`
	} `json:"transfers"`
}

func (wallet *Wallet) IncomingTransfers(id string, params IncomingTransfersParams) (result IncomingTransfersResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:incoming_transfers:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "incoming_transfers",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:incoming_transfers", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:incoming_transfers:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:incoming_transfers:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
