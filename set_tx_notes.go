package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type SetTxNotesParams struct {
	TxIDs []string `json:"tx_ids"`
	Notes []string `json:"notes"`
}

func (wallet *Wallet) SetTxNotes(id string, params SetTxNotesParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:set_tx_notes:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "set_tx_notes",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:set_tx_notes", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:set_tx_notes:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:set_tx_notes:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("set tx notes failed")
	}
	aLog.Success("xmrLib:set_tx_notes:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
