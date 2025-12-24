package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type SetAttributeParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (wallet *Wallet) SetAttribute(id string, params SetAttributeParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:set_attribute:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "set_attribute",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:set_attribute", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:set_attribute:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:set_attribute:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("set attribute failed")
	}
	aLog.Success("xmrLib:set_attribute:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
