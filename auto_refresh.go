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
	if DebugLevel >= DebugLevel3 {
		aLog.Debug(aLog.Log{Sender: "xmrLib:auto_refresh", Message: fmt.Sprintf("Params: %v", params)})
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "auto_refresh",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		if DebugLevel >= DebugLevel1 {
			aLog.Error(aLog.Log{Sender: "xmrLib:auto_refresh", Message: fmt.Errorf("error: %v", err)})
		}
		return err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return err
		} else {
			if len(jrpcRes.Result) == 0 {
				return nil
			} else {
				return errors.New("auto refresh failed")
			}
		}
	}
}

//working on adding propper debug logging
