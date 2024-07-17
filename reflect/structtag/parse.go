package structtag

import (
	reflecti "github.com/hopeio/utils/reflect/converter"
	"reflect"
	"strings"
)

// ParseTagSetting default sep ;
func ParseTagSettingToMap(tagStr string, sep string) map[string]string {
	if tagStr == "" {
		return nil
	}
	if sep == "" {
		sep = ";"
	}
	settings := map[string]string{}
	names := strings.Split(tagStr, sep)

	for i := 0; i < len(names); i++ {
		j := i
		if len(names[j]) > 0 {
			for {
				if names[j][len(names[j])-1] == '\\' {
					i++
					names[j] = names[j][0:len(names[j])-1] + sep + names[i]
					names[i] = ""
				} else {
					break
				}
			}
		}

		values := strings.Split(names[j], ":")
		k := strings.TrimSpace(strings.ToUpper(values[0]))

		if len(values) >= 2 {
			settings[k] = strings.Join(values[1:], ":")
		} else if k != "" {
			settings[k] = "true"
		}
	}

	return settings
}

func ParseTagSettingToStruct[T any](tagStr string, sep string, metaTag string) (*T, error) {
	settings := new(T)
	err := ParseTagSettingIntoStruct(tagStr, sep, settings, metaTag)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

// ParseTagSettingInto default sep ;
func ParseTagSettingIntoStruct(tagStr string, sep string, settings any, metaTag string) error {
	tagSettings := ParseTagSettingToMap(tagStr, sep)
	if tagSettings == nil {
		return errTagNotExist
	}
	settingsValue := reflect.ValueOf(settings).Elem()
	settingsType := reflect.TypeOf(settings).Elem()
	for i := 0; i < settingsValue.NumField(); i++ {
		structField := settingsType.Field(i)
		name := structField.Name
		if metaTag != "" {
			if metatag, ok := structField.Tag.Lookup(metaTag); ok {
				name = metatag
			}
		}
		if flagtag, ok := tagSettings[strings.ToUpper(name)]; ok {
			err := reflecti.SetValueByString(settingsValue.Field(i), flagtag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
