package xmrLib

import (
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

func (wallet *Wallet) StopWallet(id string) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:stop_wallet:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "stop_wallet",
		Params:  map[string]interface{}{},
	}); err != nil {
		aLog.Error("xmrLib:stop_wallet", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:stop_wallet:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:stop_wallet:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("stop wallet failed")
	}
	aLog.Success("xmrLib:stop_wallet:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
