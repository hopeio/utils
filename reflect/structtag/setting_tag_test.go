package structtag

import (
	"reflect"
	"testing"
)

type Bar2 struct {
	Field1 int    `type:"\\w"`
	Field2 string `mock:"example:1;type:\\w;test:\""`
}

func TestSettingTag(t *testing.T) {
	var bar Bar2
	typ := reflect.TypeOf(bar)

	t.Log(ParseSettingTagToMap(typ.Field(1).Tag.Get("mock"), ';'))
}
