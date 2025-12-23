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
		aLog.Error(aLog.Log{Sender: "xmrLib:bytesToMap", Message: "Error: " + err.Error()})
		return nil
	} else {
		var mymap map[string]interface{}
		if err = json.Unmarshal(data, &mymap); err != nil {
			aLog.Error(aLog.Log{Sender: "xmrLib:bytesToMap", Message: string(data)})
			aLog.Error(aLog.Log{Sender: "xmrLib:bytesToMap", Message: "Error: " + err.Error()})
			return nil
		}
		if DebugLevel >= DebugLevel1 {
			aLog.Debug(aLog.Log{Sender: "xmrLib:bytesToMap", Message: mymap})
		}
		return mymap
	}
}

// bytesToJRPCResult converts a JSON-encoded byte slice into a JRPCResult.
// Returns nil if an error occurs during unmarshalling or if the input error is not nil.
func bytesToJRPCResult(data []byte) (result *jrpcLib.JRPCResult, err error) {
	err = json.Unmarshal(data, &result)
	if err != nil {
		aLog.Error(aLog.Log{Sender: "xmrLib:bytesToJRPCResult", Message: "Data : " + string(data)})
		aLog.Error(aLog.Log{Sender: "xmrLib:bytesToJRPCResult", Message: "Error: " + err.Error()})
	}
	if DebugLevel >= DebugLevel1 {
		aLog.Debug(aLog.Log{Sender: "xmrLib:bytesToJRPCResult", Message: "Result : " + fmt.Sprint(result)})
	}
	return result, err
}

// mapToStruct converts a map[string]interface{} into a struct.
func mapToStruct(data map[string]interface{}, v interface{}) error {
	if bytes, err := json.Marshal(data); err != nil {
		aLog.Error(aLog.Log{Sender: "xmrLib:mapToStruct", Message: "Data : " + fmt.Sprint(data)})
		aLog.Error(aLog.Log{Sender: "xmrLib:mapToStruct", Message: "Error: " + err.Error()})
		return err
	} else {
		if DebugLevel >= DebugLevel1 {
			aLog.Debug(aLog.Log{Sender: "xmrLib:mapToStruct", Message: "Bytes : " + string(bytes)})
		}
		return json.Unmarshal(bytes, v)
	}
}

// Request sends a JSON-RPC request to the daemon and returns the response.
func (wallet *Wallet) Request(jrpc *jrpcLib.JRPC) (result *jrpcLib.JRPCResult, err error) {
	if res, err := wallet.Call(jrpc); err != nil {
		aLog.Error(aLog.Log{Sender: "xmrLib:Request", Message: "Error: " + err.Error()})
		return nil, err
	} else {
		if DebugLevel >= DebugLevel1 {
			aLog.Debug(aLog.Log{Sender: "xmrLib:Request", Message: "Result : " + fmt.Sprint(res)})
		}
		if checkStatus(res) != nil {
			aLog.Debug(aLog.Log{Sender: "xmrLib:Request", Message: fmt.Sprintf("Status: %d Error: %s", res.StatusCode, res.Body)})
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
