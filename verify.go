package xmrLib

import (
	"encoding/json"

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
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "verify",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
			return result, err
		} else {
			return result, mapToStruct(jrpcRes.Result, &result)
		}
	}
}
