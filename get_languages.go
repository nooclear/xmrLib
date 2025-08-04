package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

// GetLanguages retrieves a list of supported languages for the wallet.
// Executes a JSON-RPC call with the specified ID and returns the result as a JSON-encoded byte slice or an error.
func (wallet *Wallet) GetLanguages(id string) ([]byte, error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_languages",
			Params:  nil,
		}); err != nil {
		return nil, err
	} else {
		defer func() {
			err = res.Body.Close() // need to find a way to properly handle these errors
		}()
		return checkStatus(res)
	}
}
