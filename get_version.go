package xmrLib

import (
	"github.com/nooclear/jrpcLib"
)

// GetVersion retrieves the wallet's version via JSON-RPC using the provided ID and returns it as a byte slice or an error.
func (wallet *Wallet) GetVersion(id string) ([]byte, error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_version",
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
