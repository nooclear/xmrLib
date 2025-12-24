package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type AutoRefreshParams struct {
	Enabled bool   `json:"enabled"`
	Period  uint64 `json:"period"`
}

func (wallet *Wallet) AutoRefresh(id string, params AutoRefreshParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:auto_refresh:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "auto_refresh",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:auto_refresh", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:auto_refresh:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:auto_refresh:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("auto refresh failed")
	}
	aLog.Success("xmrLib:auto_refresh:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
