package util

import "encoding/json"

// ToJSONString converts an object to a JSON string.
func ToJSONString(obj any) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
