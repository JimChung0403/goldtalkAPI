package util

import (
	"encoding/json"

	"github.com/json-iterator/go"
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func JsonString(v interface{}) string {
	Json, _ := Json.Marshal(v)
	return string(Json)
}

func JsonUnmarshalFromString(jsonStr string, v interface{}) error {
	err := Json.UnmarshalFromString(jsonStr, v)
	if err != nil {
		return json.Unmarshal([]byte(jsonStr), v)
	}
	return err
}
