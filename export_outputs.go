package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type ExportOutputsParams struct {
	All bool `json:"all"`
}

type ExportOutputsResult struct {
	OutputsDataHex string `json:"outputs_data_hex"`
}

func (wallet *Wallet) ExportOutputs(id string, params ExportOutputsParams) (result ExportOutputsResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "export_outputs",
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
