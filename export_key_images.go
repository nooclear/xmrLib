package xmrLib

import (
	"encoding/json"

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
	if res, err := wallet.Call(&jrpcLib.JRPC{
		Version: JRPCVersion,
		ID:      id,
		Method:  "export_key_images",
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
