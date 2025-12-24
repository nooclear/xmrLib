package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type ExportKeyImagesParams struct {
	AccountIndex uint64 `json:"account_index"`
}

type ExportKeyImagesResult struct {
	Offset          uint64 `json:"offset"`
	SignedKeyImages []struct {
		KeyImage  string `json:"key_image"`
		Signature string `json:"signature"`
	} `json:"signed_key_images"`
}

func (wallet *Wallet) ExportKeyImages(id string, params ExportKeyImagesParams) (result ExportKeyImagesResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:export_key_images:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "export_key_images",
		Params:  bytesToMap(json.Marshal(params)),
	}); err != nil {
		aLog.Error("xmrLib:export_key_images", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:export_key_images:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:export_key_images:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
