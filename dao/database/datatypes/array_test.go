package datatypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	data := `{{{{1,2},{3,4}},{{5,6},{7,8}}}}`
	arr := Array[Array[Array[Array[int]]]]{}
	err := arr.Scan(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(arr)
	t.Log(arr.Value())
}

func TestJSONArray(t *testing.T) {
	data := `{"{\"name\": \"key1\", \"value\": [{\"subKey1\": \"value1\", \"subKey2\": \"value2\"}]}","{\"name\": \"key2\", \"value\": \"\"}","{\"name\": \"key3\", \"value\": \"\"}","{\"name\": \"key4\", \"value\": [{\"subKey3\": \"value3\", \"subKey4\": [\"value4\"]}]}","{\"name\": \"key5\", \"value\": null}","{\"name\": \"key6\", \"value\": []}"}`
	var jat JsonArray
	err := jat.Scan([]byte(data))
	if err != nil {
		t.Error(err)
	}
	t.Log(jat)
	assert.Equal(t, JsonArray{map[string]interface{}{"name": "key1", "value": []interface{}{map[string]interface{}{"subKey1": "value1", "subKey2": "value2"}}}, map[string]interface{}{"name": "key2", "value": ""}, map[string]interface{}{"name": "key3", "value": ""}, map[string]interface{}{"name": "key4", "value": []interface{}{map[string]interface{}{"subKey3": "value3", "subKey4": []interface{}{"value4"}}}}, map[string]interface{}{"name": "key5", "value": interface{}(nil)}}, jat)
}
