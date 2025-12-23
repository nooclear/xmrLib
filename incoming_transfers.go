package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type IncomingTransfersParams struct {
	TransferType   Transfer `json:"transfer_type"`
	AccountIndex   uint64   `json:"account_index"`
	SubaddrIndices []uint64 `json:"subaddr_indices"`
}

type Transfer string

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
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "incoming_transfers",
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
