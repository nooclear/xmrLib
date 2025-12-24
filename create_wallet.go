package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type CreateWalletParams struct {
	Filename string `json:"filename"`
	Password string `json:"password"`
	Language string `json:"language"`
}

func (wallet *Wallet) CreateWallet(id string, params CreateWalletParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:create_wallet:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "create_wallet",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:create_wallet", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:create_wallet:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:create_wallet:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("create wallet failed")
	}
	aLog.Success("xmrLib:create_wallet:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
