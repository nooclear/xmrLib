package xmrLib

import (
	"encoding/json"

	"github.com/nooclear/jrpcLib"
)

/*
account_index - unsigned int; Return balance for this account.
address_indices - array of unsigned int; (Optional) Return balance detail for those subaddresses.
all_accounts - boolean; (Defaults to false)
strict - boolean; (Defaults to false) all changes go to 0-th subaddress (in the current subaddress account)
*/

// balanceParams defines the parameters for querying wallet balance via JSON-RPC calls.
type BalanceParams struct {
	AccountIndex   uint64   `json:"account_index"`
	AddressIndices []uint64 `json:"address_indices"`
	AllAccounts    bool     `json:"all_accounts"`
	Strict         bool     `json:"strict"`
}

type BalanceResult struct {
	Balance              uint64 `json:"balance"`
	BlocksToUnlock       uint64 `json:"blocks_to_unlock"`
	MultisigImportNeeded bool   `json:"multisig_import_needed"`
	TimeToUnlock         uint64 `json:"time_to_unlock"`
	UnlockedBalance      uint64 `json:"unlocked_balance"`
}

// GetBalance retrieves the balance details for a wallet based on the given parameters.
// It uses a JSON-RPC call with specified ID and parameters.
// Returns the balance as a JSON-encoded byte slice, or an error if the request fails.
func (wallet *Wallet) GetBalance(id string, params BalanceParams) (result BalanceResult, err error) {
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_balance",
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
