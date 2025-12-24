package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type LabelAccountParams struct {
	AccountIndex uint64 `json:"account_index"`
	Label        string `json:"label"`
}

func (wallet *Wallet) LabelAccount(id string, params LabelAccountParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:label_account:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "label_account",
			Params:  bytesToMap(json.Marshal(params)),
		}); err != nil {
		aLog.Error("xmrLib:label_account", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:label_account:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:label_account:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("label account failed")
	}
	aLog.Success("xmrLib:label_account:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
