package jsonutils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ToJsonString(value interface{}) string {
	return string(ToJsonBytes(value))
}

func ToJsonBytes(value interface{}) []byte {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return json
}

func DecodeIoReader[T any](ioReader io.ReadCloser) *T {
	var input T
	err := json.NewDecoder(ioReader).Decode(&input)
	if err != nil {
		panic(err)
	}

	return &input
}

func EncodeJson[T any](w http.ResponseWriter, toEncode T) {
	err := json.NewEncoder(w).Encode(toEncode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
