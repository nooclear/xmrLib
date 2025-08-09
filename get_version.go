package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

type VersionResult struct {
	Release bool   `json:"release"`
	Version uint64 `json:"version"`
}

// GetVersion retrieves the wallet's version via JSON-RPC using the provided ID and returns it as a byte slice or an error.
func (wallet *Wallet) GetVersion(id string) (result VersionResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_version",
			Params:  nil,
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
