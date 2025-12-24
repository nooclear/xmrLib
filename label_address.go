package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type LabelAddressParams struct {
	Index struct {
		Major uint64 `json:"major"`
		Minor uint64 `json:"minor"`
	} `json:"index"`
	Label string `json:"label"`
}

func (wallet *Wallet) LabelAddress(id string, params LabelAddressParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:label_address:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "label_address",
			Params:  bytesToMap(json.Marshal(params)),
		}); err != nil {
		aLog.Error("xmrLib:label_address", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:label_address:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:label_address:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("label address failed")
	}
	aLog.Success("xmrLib:label_address:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
