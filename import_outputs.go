package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type ImportOutputsParams struct {
	OutputsDataHex string `json:"outputs_data_hex"`
}

type ImportOutputsResult struct {
	NumImported uint64 `json:"num_imported"`
}

func (wallet *Wallet) ImportOutputs(id string, params ImportOutputsParams) (result ImportOutputsResult, err error) {
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "import_outputs",
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
