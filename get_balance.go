package xmrLib

import (
	"encoding/json"
	"errors"
	"github.com/nooclear/jrpcLib"
	"io"
)

/*
account_index - unsigned int; Return balance for this account.
address_indices - array of unsigned int; (Optional) Return balance detail for those subaddresses.
all_accounts - boolean; (Defaults to false)
strict - boolean; (Defaults to false) all changes go to 0-th subaddress (in the current subaddress account)
*/

// balanceParams defines the parameters for querying wallet balance via JSON-RPC calls.
type balanceParams struct {
	AccountIndex   uint64   `json:"account_index"`
	AddressIndices []uint64 `json:"address_indices"`
	AllAccounts    bool     `json:"all_accounts"`
	Strict         bool     `json:"strict"`
}

// GetBalance retrieves the balance details for a wallet based on the given parameters.
// It uses a JSON-RPC call with specified ID and parameters.
// Returns the balance as a JSON-encoded byte slice, or an error if the request fails.
func (wallet *Wallet) GetBalance(id string, params balanceParams) ([]byte, error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_balance",
			Params:  convertToMap(json.Marshal(params)),
		}); err != nil {
		return nil, err
	} else {
		defer func() {
			err = res.Body.Close() // need to find a proper way to handle these errors
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
