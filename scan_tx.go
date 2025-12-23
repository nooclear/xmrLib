package xmrLib

import (
	"encoding/json"
	"errors"

	"github.com/nooclear/jrpcLib"
)

type ScanTxParams struct {
	TxIDs []string `json:"txids"`
}

func (wallet *Wallet) ScanTx(id string, params ScanTxParams) (err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "scan_tx",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		return err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return err
		} else {
			if len(jrpcRes.Result) == 0 {
				return nil
			} else {
				return errors.New("scan tx failed")
			}
		}
	}
}
