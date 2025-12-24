package xmrLib

import (
	"encoding/json"
	"errors"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ScanTxParams struct {
	TxIDs []string `json:"txids"`
}

func (wallet *Wallet) ScanTx(id string, params ScanTxParams) (err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:scan_tx:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "scan_tx",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:scan_tx", fmt.Sprintf("error: %v", err))
		return err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:scan_tx:jrpcRes", fmt.Sprintf("error: %v", err))
		return err
	} else if len(jrpcRes.Result) != 0 {
		aLog.Error("xmrLib:scan_tx:jrpcRes.Result", fmt.Sprintf("result: %v", jrpcRes.Result))
		return errors.New("scan tx failed")
	}
	aLog.Success("xmrLib:scan_tx:success", fmt.Sprintf("wallet: %v", wallet))
	return nil
}
