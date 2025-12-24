package xmrLib

import (
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

func (wallet *Wallet) Store(id string) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:store:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "store",
		Params:  map[string]interface{}{},
	}); err != nil {
		aLog.Error("xmrLib:store", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:store:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:store:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("store failed")
	}
	aLog.Success("xmrLib:store:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
