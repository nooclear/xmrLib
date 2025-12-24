package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type LanguagesResult struct {
	Languages      []string `json:"languages"`
	LanguagesLocal []string `json:"languages_local"`
}

// GetLanguages retrieves a list of supported languages for the wallet.
// Executes a JSON-RPC call with the specified ID and returns the result as a JSON-encoded byte slice or an error.
func (wallet *Wallet) GetLanguages(id string) (result LanguagesResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_languages:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_languages",
			Params:  nil,
		}); err != nil {
		aLog.Error("xmrLib:get_languages", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_languages:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_languages:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}

// todo
func convertToLanguagesResult(data map[string]interface{}) (result LanguagesResult, err error) {
	if bytes, err := json.Marshal(data); err != nil {
		return result, err
	} else {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
}
