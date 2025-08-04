package xmrLib

import (
	"encoding/json"
	"fmt"
)

// convertToMap converts a JSON-encoded byte slice into a map[string]interface{}.
// Returns nil if an error occurs during unmarshalling or if the input error is not nil.
func convertToMap(data []byte, err error) map[string]interface{} {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil
	} else {
		var mymap map[string]interface{}
		if _ = json.Unmarshal(data, &mymap); err != nil {
			fmt.Println("Error: " + err.Error())
			return nil
		}
		return mymap
	}
}
