package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ChangeWalletPasswordParams struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (wallet *Wallet) ChangeWalletPassword(id string, params ChangeWalletPasswordParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:change_wallet_password:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "change_wallet_password",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:change_wallet_password", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:auto_refresh:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:auto_refresh:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("change wallet password failed")
	}
	aLog.Success("xmrLib:change_wallet_password:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
