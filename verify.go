package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type VerifyParams struct {
	Data      string `json:"data"`
	Address   string `json:"address"`
	Signature string `json:"signature"`
}
type VerifyResult struct {
	Good bool `json:"good"`
}

func (wallet *Wallet) Verify(id string, params VerifyParams) (result VerifyResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:verify:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "verify",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:verify", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:verify:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:verify:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
