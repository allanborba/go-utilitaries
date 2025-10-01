package jsonutils

import "encoding/json"

func ToJsonBytes(value interface{}) []byte {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return json
}
