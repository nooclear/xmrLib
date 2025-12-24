package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type TagAccountsParams struct {
	Tag      string   `json:"tags"`
	Accounts []uint64 `json:"accounts"`
}

func (wallet *Wallet) TagAccounts(id string, params TagAccountsParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:tag_accounts:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "tag_accounts",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:tag_accounts", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:tag_accounts:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:tag_accounts:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("tag accounts failed")
	}
	aLog.Success("xmrLib:tag_accounts:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
