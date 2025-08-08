package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type VersionResult struct {
	Release bool   `json:"release"`
	Version uint64 `json:"version"`
}

// GetVersion retrieves the wallet's version via JSON-RPC using the provided ID and returns it as a byte slice or an error.
func (wallet *Wallet) GetVersion(id string) (verRes VersionResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_version",
			Params:  nil,
		}); err != nil {
		return verRes, err
	} else {
		if jrpcRes, err := convertToJRPCResult(res.Body); err != nil {
			return verRes, err
		} else {
			return convertToVersionResult(jrpcRes.Result)
		}
	}
}

func convertToVersionResult(data map[string]interface{}) (result VersionResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
