package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

/*
tag - string; (Optional) Tag for filtering accounts.
regex - boolean; (Optional) allow regular expression filters if set to true (Defaults to false).
strict_balances - boolean; (Optional) when true, balance only considers the blockchain, when false it considers both the blockchain and some recent actions, such as a recently created transaction which spent some outputs, which isn't yet mined.

Outputs:
subaddress_accounts - array of subaddress account information:
    account_index - unsigned int; Index of the account.
    balance - unsigned int; Balance of the account (locked or unlocked).
    base_address - string; Base64 representation of the first subaddress in the account.
    label - string; (Optional) Label of the account.
    tag - string; (Optional) Tag for filtering accounts.
    unlocked_balance - unsigned int; Unlocked balance for the account.
total_balance - unsigned int; Total balance of the selected accounts (locked or unlocked).
total_unlocked_balance - unsigned int; Total unlocked balance of the selected accounts.
//
*/

// accountsParams defines parameters and details of accounts, including tag, balances, subaccount info, and balance types.
type accountsParams struct {
	Tag            string `json:"tag"`
	Regex          bool   `json:"regex"`
	StrictBalances bool   `json:"strict_balances"`
}

type AccountsResult struct {
	SubaddressAccounts []struct {
		AccountIndex    uint64 `json:"account_index"`
		Balance         uint64 `json:"balance"`
		BaseAddress     string `json:"base_address"`
		Label           string `json:"label"`
		Tag             string `json:"tag"`
		UnlockedBalance uint64 `json:"unlocked_balance"`
	} `json:"subaddress_accounts"`
}

// GetAccounts retrieves account details for a wallet using the specified ID via JSON-RPC.
// Returns the account information as a JSON-encoded byte slice or an error if the request fails.
func (wallet *Wallet) GetAccounts(id string, params accountsParams) (result AccountsResult, err error) {
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:get_accounts:start", fmt.Sprintf("wallet: %v", wallet))
	}
	if res, err := wallet.Call(
		&jrpcLib.JRPC{
			Version: JRPCVersion,
			ID:      id,
			Method:  "get_accounts",
			Params:  bytesToMap(json.Marshal(params)),
		}); err != nil {
		aLog.Error("xmrLib:get_accounts", fmt.Sprintf("error: %v", err))
		return result, err
	} else if jrpcRes, err := bytesToJRPCResult(res.Body); err != nil {
		aLog.Error("xmrLib:get_accounts:jrpcRes", fmt.Sprintf("error: %v", err))
		return result, err
	} else {
		aLog.Success("xmrLib:get_accounts:success", fmt.Sprintf("wallet: %v", wallet))
		return result, mapToStruct(jrpcRes.Result, &result)
	}
}
