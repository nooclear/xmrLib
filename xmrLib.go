package xmrLib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var JRPCVersion = "2.0"

// convertToMap converts a JSON-encoded byte slice into a map[string]interface{}.
// Returns nil if an error occurs during unmarshalling or if the input error is not nil.
func convertToMap(data []byte, err error) map[string]interface{} {
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

// responseToStruct converts JSON-encoded byte data into a WalletResponse struct and returns it or an error if unmarshalling fails.
func responseToStruct(data []byte) (*WalletResponse, error) {
	res := WalletResponse{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// checkStatus processes an HTTP response, returning the body as a byte slice if the status is 200, or an error otherwise.
func checkStatus(response *http.Response) ([]byte, error) {
	switch response.StatusCode {
	case 200:
		if data, err := io.ReadAll(response.Body); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	default:
		return nil, fmt.Errorf("unhandled status code: %d %s", response.StatusCode, response.Status)
	}
}
