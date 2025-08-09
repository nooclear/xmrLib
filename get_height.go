package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

type HeightResult struct {
	Height uint64 `json:"height"`
}

// GetHeight retrieves the current blockchain height via JSON-RPC using the provided ID.
// Returns the height as a JSON-encoded byte slice, or an error if the request fails.
func (wallet *Wallet) GetHeight(id string) (result HeightResult, err error) {
	if httpRes, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_height",
			Params:  nil,
		}); err != nil {
		return result, err
	} else {
		if jrpcRes, err := bytesToJRPCResult(httpRes.Body); err != nil {
			return result, err
		} else {
			return result, mapToStruct(jrpcRes.Result, &result)
		}
	}
}

// convertToHeightResult converts a map of string to interface{} into a HeightResult struct.
// Returns the converted result or an error if the conversion fails.
