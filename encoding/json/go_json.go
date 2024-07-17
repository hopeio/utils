//go:build go_json

package json

import json "github.com/goccy/go-json"

var (
	Marshal = json.Marshal

	Unmarshal = json.Unmarshal

	MarshalIndent = json.MarshalIndent

	NewDecoder = json.NewDecoder

	NewEncoder = json.NewEncoder
)

func MarshalToString(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return strings.ToString(data), nil
}

func UnmarshalFromString(str string, v any) error {
	return json.Unmarshal(strings.ToBytes(str), v)
}
