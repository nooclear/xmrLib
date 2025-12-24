package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ImportOutputsParams struct {
	OutputsDataHex string `json:"outputs_data_hex"`
}

type ImportOutputsResult struct {
	NumImported uint64 `json:"num_imported"`
}

func (wallet *Wallet) ImportOutputs(id string, params ImportOutputsParams) (result ImportOutputsResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:import_outputs:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "import_outputs",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:import_outputs", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:import_outputs:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:import_outputs:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
