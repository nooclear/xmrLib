package xmrLib

import (
	"encoding/json"
	"github.com/nooclear/jrpcLib"
)

type HeightResult struct {
	Height uint64 `json:"height"`
}

// GetHeight retrieves the current blockchain height via JSON-RPC using the provided ID.
// Returns the height as a JSON-encoded byte slice, or an error if the request fails.
func (wallet *Wallet) GetHeight(id string) (heightResult HeightResult, err error) {
	if httpRes, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_height",
			Params:  nil,
		}); err != nil {
		return heightResult, err
	} else {
		if jrpcRes, err := convertToJRPCResult(httpRes.Body); err != nil {
			return heightResult, err
		} else {
			return convertToHeightResult(jrpcRes.Result)
		}
	}
}

// convertToHeightResult converts a map of string to interface{} into a HeightResult struct.
// Returns the converted result or an error if the conversion fails.
func convertToHeightResult(data map[string]interface{}) (result HeightResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
