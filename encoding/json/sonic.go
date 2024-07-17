//go:build sonic && amd64

package json

import "github.com/bytedance/sonic"

var (
	json = sonic.ConfigStd

	Marshal = json.Marshal

	Unmarshal = json.Unmarshal

	MarshalIndent = json.MarshalIndent

	NewDecoder = json.NewDecoder

	NewEncoder = json.NewEncoder

	MarshalToString     = json.MarshalToString
	UnmarshalFromString = json.UnmarshalFromString
)

func MarshalReader(v interface{}) (io.Reader, error) {
	data, err := sonic.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
