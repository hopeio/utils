package common

import (
	"encoding/json"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

type Foo struct {
	Level    zapcore.Level
	Duration time.Duration
}

func TestUnmarshal(t *testing.T) {
	data := []byte(`Level="debug"
	Duration="100ms"`)
	var foo Foo
	err := toml.Unmarshal(data, &foo)
	if err != nil {
		t.Error(err)
	}
	t.Log(foo.Level)
	t.Log(foo.Duration)
}

type Bar struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
	Foo    `json:"foo" toml:"foo"`
}

func TestAnonymous(t *testing.T) {
	var bar Bar
	data, err := json.Marshal(bar)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
	data, err = toml.Marshal(bar)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
	var bar2 Bar
	data = []byte(`{"field1":"2","field2":1,"Level":"debug","Duration":2}`)
	err = json.Unmarshal(data, &bar2)
	if err != nil {
		t.Error(err)
	}
	t.Log(bar2)
	var bar3 Bar
	data = []byte(`{"field1":"3","field2":3,"foo":{"Level":"debug","Duration":"3s"}}`)
	err = json.Unmarshal(data, &bar3)
	if err != nil {
		t.Error(err)
	}
	t.Log(bar3)
}
