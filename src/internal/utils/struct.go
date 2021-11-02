package utils

import (
	"encoding/json"
)

func Bind(dict map[string]interface{}, obj interface{}) error {
	jsonbody, err := json.Marshal(dict)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonbody, obj); err != nil {
		return err
	}

	return nil
}
