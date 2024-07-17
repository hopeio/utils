package sjson

import "github.com/tidwall/sjson"

func Set(json, path string, value interface{}) (string, error) {
	return sjson.SetOptions(json, path, value, nil)
}
