package xmrLib

import (
	"errors"
	"github.com/nooclear/jrpcLib"
	"io"
)

// GetVersion retrieves the wallet's version via JSON-RPC using the provided ID and returns it as a byte slice or an error.
func (wallet *Wallet) GetVersion(id string) ([]byte, error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: "2.0",
			ID:      id,
			Method:  "get_version",
			Params:  nil,
		}); err != nil {
		return nil, err
	} else {
		defer func() {
			err = res.Body.Close() // need to find a way to properly handle these errors
		}()
		switch res.StatusCode {
		case 200:
			if data, err := io.ReadAll(res.Body); err != nil {
				return nil, err
			} else {
				return data, nil
			}
		default:
			return nil, errors.New(res.Status)
		}
	}
}
