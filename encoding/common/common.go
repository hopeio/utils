package common

import (
	"encoding/json"
	"errors"
	"github.com/hopeio/utils/encoding"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"strings"
)

func Unmarshal(format encoding.Format, data []byte, v any) (encoding.Format, error) {
	switch {
	case format == encoding.Yaml || format == encoding.Yml:
		return format, yaml.Unmarshal(data, v)
	case format == encoding.Toml:
		return format, toml.Unmarshal(data, v)
	case format == encoding.Json:
		return format, json.Unmarshal(data, v)
	default:
		if err := toml.Unmarshal(data, v); err == nil {
			return encoding.Toml, nil
		}
		if err := json.Unmarshal(data, v); err == nil {
			return encoding.Json, nil
		} else if strings.Contains(err.Error(), "json: unknown field") {
			return encoding.Json, err
		}

		yamlError := yaml.Unmarshal(data, v)

		if yamlError == nil {
			return encoding.Yaml, nil
		} else if yErr, ok := yamlError.(*yaml.TypeError); ok {
			return encoding.Yaml, yErr
		}

		return "", errors.New("failed to decode")
	}
}

func Marshal(format encoding.Format, v any) ([]byte, error) {
	switch {
	case format == encoding.Yaml || format == encoding.Yml:
		return yaml.Marshal(v)
	case format == encoding.Toml:
		return toml.Marshal(v)
	case format == encoding.Json:
		return json.Marshal(v)
	default:
		return toml.Marshal(v)
	}
}
