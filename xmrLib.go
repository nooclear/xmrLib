package xmrLib

import (
	"encoding/json"
	"fmt"

	aLog "github.com/nooclear/AdvancedLogging"
	"github.com/nooclear/jrpcLib"
)

type Level int

var DebugLevel Level

const (
	DebugLevel0 Level = iota
	DebugLevel1
	DebugLevel2
	DebugLevel3
)

var JRPCVersion = "2.0"

// bytesToMap converts a JSON-encoded byte slice into a map[string]interface{}.
// Returns nil if an error occurs during unmarshalling or if the input error is not nil.
func bytesToMap(data []byte, err error) map[string]interface{} {
	if err != nil {
		aLog.Error("xmrLib:bytes_to_map", fmt.Sprintf("error: %v", err))
		return nil
	} else {
		var mymap map[string]interface{}
		if err = json.Unmarshal(data, &mymap); err != nil {
			aLog.Error("xmrLib:bytes_to_map:unmarshal", fmt.Sprintf("Data: %v", string(data)))
			aLog.Error("xmrLib:bytesToMap", fmt.Sprintf("Error: %v", err))
			return nil
		}
		if DebugLevel >= DebugLevel1 {
			aLog.Debug("xmrLib:bytes_to_map", fmt.Sprintf("map: %v", mymap))
		}
		return mymap
	}
}

// bytesToJRPCResult converts a JSON-encoded byte slice into a JRPCResult.
// Returns nil if an error occurs during unmarshalling or if the input error is not nil.
func bytesToJRPCResult(data []byte) (result *jrpcLib.JRPCResult, err error) {
	err = json.Unmarshal(data, &result)
	if err != nil {
		aLog.Error("xmrLib:bytes_to_jrpc_result", fmt.Sprintf("Data: %v", string(data)))
		aLog.Error("xmrLib:bytes_to_jrpc_result", fmt.Sprintf("Error: %v", err))
	}
	if DebugLevel >= DebugLevel1 {
		aLog.Debug("xmrLib:bytes_to_jrpc_result", fmt.Sprintf("Result: %v", fmt.Sprint(result)))
	}
	return result, err
}

// mapToStruct converts a map[string]interface{} into a struct.
func mapToStruct(data map[string]interface{}, v interface{}) error {
	if bytes, err := json.Marshal(data); err != nil {
		aLog.Error("xmrLib:map_to_struct", fmt.Sprintf("Data: %v", data))
		aLog.Error("xmrLib:map_to_struct", fmt.Sprintf("Error: %v", err))
		return err
	} else {
		if DebugLevel >= DebugLevel1 {
			aLog.Debug("xmrLib:map_to_struct", fmt.Sprintf("Bytes: %v", string(bytes)))
		}
		return json.Unmarshal(bytes, v)
	}
}

// Request sends a JSON-RPC request to the daemon and returns the response.
func (wallet *Wallet) Request(jrpc *jrpcLib.JRPC) (result *jrpcLib.JRPCResult, err error) {
	if res, err := wallet.Call(jrpc); err != nil {
		aLog.Error("xmrLib:request", fmt.Sprintf("Error: %v", err))
		return nil, err
	} else {
		if DebugLevel >= DebugLevel1 {
			aLog.Debug("xmrLib:request", fmt.Sprintf("Result: %v", fmt.Sprint(res)))
		}
		if checkStatus(res) != nil {
			aLog.Notify("xmrLib:request", fmt.Sprintf("Status: %d\tError: %s", res.StatusCode, res.Body))
		}
		return bytesToJRPCResult(res.Body)
	}
}

// checkStatus validates the HTTP response's status code and returns an error if it's not 200.
func checkStatus(response jrpcLib.HttpResponse) error {
	if response.StatusCode != 200 {
		return fmt.Errorf("HTTP status code: %d", response.StatusCode)
	}
	return nil
}
