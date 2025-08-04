package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

// GetHeight retrieves the current blockchain height via JSON-RPC using the provided ID.
// Returns the height as a JSON-encoded byte slice, or an error if the request fails.
func (wallet *Wallet) GetHeight(id string) ([]byte, error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_height",
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
