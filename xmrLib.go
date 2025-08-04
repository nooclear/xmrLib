package xmrLib

import (
	"encoding/json"
	"fmt"
)

func convert(data []byte, err error) map[string]interface{} {
	if err == nil {
		var mymap map[string]interface{}
		if _ = json.Unmarshal(data, &mymap); err != nil {
			fmt.Println("Error: " + err.Error())
			return nil
		}
		return mymap
	} else {
		fmt.Println("Error: " + err.Error())
		return nil
	}
}
