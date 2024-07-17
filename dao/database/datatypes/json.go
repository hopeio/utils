package datatypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JsonB []byte

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 JsonB
func (j *JsonB) Scan(value interface{}) error {
	switch bytes := value.(type) {
	case []byte:
		*j = append((*j)[0:0], bytes...)
		return nil
	case string:
		*j = JsonB(bytes)
		return nil
	default:
		return errors.New(fmt.Sprint("failed to scan JsonB value:", value))
	}

}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JsonB) Value() (driver.Value, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

func (*JsonB) GormDataType() string {
	return "jsonb"
}

type Json map[string]any

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Json
func (j *Json) Scan(value interface{}) error {
	switch bytes := value.(type) {
	case []byte:
		*j = map[string]any{}
		return json.Unmarshal(bytes, j)
	case string:
		*j = map[string]any{}
		return json.Unmarshal([]byte(bytes), j)
	default:
		return errors.New(fmt.Sprint("failed to scan Value value:", value))
	}
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j Json) Value() (driver.Value, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return json.Marshal(j)
}

func (Json) GormDataType() string {
	return "jsonb"
}

type JsonArray []map[string]any

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Json
func (j *JsonArray) Scan(value interface{}) error {
	switch bytes := value.(type) {
	case []byte:
		*j = make([]map[string]any, 0)
		return json.Unmarshal(bytes, j)
	case string:
		*j = make([]map[string]any, 0)
		return json.Unmarshal([]byte(bytes), j)
	default:
		return errors.New(fmt.Sprint("failed to scan JsonArray value:", value))
	}
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JsonArray) Value() (driver.Value, error) {
	if j == nil {
		return []byte("null"), nil
	}
	if len(j) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(j)
}

func (*JsonArray) GormDataType() string {
	return "jsonb"
}

type JsonStr string

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 JsonStr
func (j *JsonStr) Scan(value interface{}) error {
	switch bytes := value.(type) {
	case []byte:
		*j = JsonStr(bytes)
		return nil
	case string:
		*j = JsonStr(bytes)
		return nil
	default:
		return errors.New(fmt.Sprint("failed to scan JsonStr value:", value))
	}
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JsonStr) Value() (driver.Value, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return string(j), nil
}

func (*JsonStr) GormDataType() string {
	return "jsonb"
}

type JsonT[T any] struct {
	Data *T
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Json
func (j *JsonT[T]) Scan(value interface{}) error {
	switch bytes := value.(type) {
	case []byte:
		j.Data = new(T)
		return json.Unmarshal(bytes, j)
	case string:
		j.Data = new(T)
		return json.Unmarshal([]byte(bytes), j)
	default:
		return errors.New(fmt.Sprint("failed to scan Value value:", value))
	}
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JsonT[T]) Value() (driver.Value, error) {
	if j.Data == nil {
		return []byte("null"), nil
	}
	return json.Marshal(j)
}

func (*JsonT[T]) GormDataType() string {
	return "jsonb"
}

type JsonTArray[T any] []T

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Json
func (j *JsonTArray[T]) Scan(value interface{}) error {
	switch bytes := value.(type) {
	case []byte:
		*j = make(JsonTArray[T], 0)
		return json.Unmarshal(bytes, j)
	case string:
		*j = make(JsonTArray[T], 0)
		return json.Unmarshal([]byte(bytes), j)
	default:
		return errors.New(fmt.Sprint("failed to scan Value value:", value))
	}
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JsonTArray[T]) Value() (driver.Value, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return json.Marshal(j)
}

func (*JsonTArray[T]) GormDataType() string {
	return "jsonb"
}
