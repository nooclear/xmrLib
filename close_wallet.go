package xmrLib

import (
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

func (wallet *Wallet) CloseWallet(id string) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:close_wallet:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "close_wallet",
		Params:  map[string]interface{}{},
	}); err != nil {
		return err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			aLog.Error("xmrLib:close_wallet", fmt.Sprintf("error: %v", err))
			return err
		} else if len(jrpcRes.Result) != 0 {
			aLog.Error("xmrLib:close_wallet:jrpcRes.Result", fmt.Sprintf("error: %v", jrpcRes.Result))
			return errors.New("close wallet failed")
		}
		aLog.Success("xmrLib:close_wallet:success", fmt.Sprintf("wallet: %v", wallet))
		return nil
	}
}
