package xmrLib

import (
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type HeightResult struct {
	Height uint64 `json:"height"`
}

// GetHeight retrieves the current blockchain height via JSON-RPC using the provided ID.
// Returns the height as a JSON-encoded byte slice, or an error if the request fails.
func (wallet *Wallet) GetHeight(id string) (result HeightResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_height:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if httpRes, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_height",
			Params:  nil,
		}); err != nil {
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(httpRes.Body); err != nil {
		aLog.Error("xmrLib:get_height:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_height:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}

// convertToHeightResult converts a map of string to interface{} into a HeightResult struct.
// Returns the converted result or an error if the conversion fails.
