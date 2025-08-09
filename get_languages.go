package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

type LanguagesResult struct {
	Languages      []string `json:"languages"`
	LanguagesLocal []string `json:"languages_local"`
}

// GetLanguages retrieves a list of supported languages for the wallet.
// Executes a JSON-RPC call with the specified ID and returns the result as a JSON-encoded byte slice or an error.
func (wallet *Wallet) GetLanguages(id string) (result LanguagesResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_languages",
			Params:  nil,
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

func convertToLanguagesResult(data map[string]interface{}) (result LanguagesResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
