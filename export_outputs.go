package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ExportOutputsParams struct {
	All bool `json:"all"`
}

type ExportOutputsResult struct {
	OutputsDataHex string `json:"outputs_data_hex"`
}

func (wallet *Wallet) ExportOutputs(id string, params ExportOutputsParams) (result ExportOutputsResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:export_outputs:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "export_outputs",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:export_outputs", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:export_outputs:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:export_outputs:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
