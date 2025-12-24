package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type QueryKeyParams struct {
	KeyType keyType `json:"key_type"`
}

type keyType string

var mnemonic keyType = "mnemonic"
var viewkey keyType = "view_key"
var spendkey keyType = "spend_key"

type QueryKeyResult struct {
	Key string `json:"key"`
}

func (wallet *Wallet) QueryKey(id string, params QueryKeyParams) (result QueryKeyResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:query_key:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "query_key",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:query_key", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:query_key:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:query_key:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
