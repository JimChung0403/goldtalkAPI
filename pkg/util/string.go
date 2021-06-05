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


func InStringSlice(val string, ss []string) bool {
	if len(val) == 0 {
		return false
	}
	for _, s := range ss {
		if val == s {
			return true
		}
	}
	return false
}