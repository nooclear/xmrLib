package xmrLib

import (
	"encoding/json"
	"fmt"

	"github.com/nooclear/jrpcLib"
)

var JRPCVersion = "2.0"

// bytesToMap converts a JSON-encoded byte slice into a map[string]interface{}.
// Returns nil if an error occurs during unmarshalling or if the input error is not nil.
func bytesToMap(data []byte, err error) map[string]interface{} {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil
	} else {
		var mymap map[string]interface{}
		if err = json.Unmarshal(data, &mymap); err != nil {
			fmt.Println("Error: " + err.Error())
			return nil
		}
		return mymap
	}
}

func bytesToJRPCResult(data []byte) (result jrpcLib.JRPCResult, err error) {
	err = json.Unmarshal(data, &result)
	return result, err
}

func mapToStruct(data map[string]interface{}, v interface{}) error {
	if bytes, err := json.Marshal(data); err != nil {
		return err
	} else {
		return json.Unmarshal(bytes, v)
	}
}
