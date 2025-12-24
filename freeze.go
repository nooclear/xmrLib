package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type FreezeParams struct {
	KeyImage string `json:"key_image"`
}

func (wallet *Wallet) Freeze(id string, params FreezeParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:freeze:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "freeze",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:freeze", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:freeze:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:freeze:jrpcRes.Result", fmt.Sprintf("error: %v", jrpcRes.Result))
		return errors.New("freeze failed")
	}
	aLog.Success("xmrLib:freeze:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
