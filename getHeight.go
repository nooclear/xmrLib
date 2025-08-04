package xmrLib

import (
	"errors"
	"fmt"
	"github.com/nooclear/jrpcLib"
	"io"
)

func (wallet *Wallet) GetHeight() ([]byte, error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: "2.0",
			ID:      "",
			Method:  "get_height",
			Params:  map[string]interface{}{},
		}); err != nil {
		return nil, err
	} else {
		defer func() {
			if err = res.Body.Close(); err != nil {
				fmt.Println(err)
			}
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
